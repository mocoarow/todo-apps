package gateway_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/gateway"
)

func newTestAuthTokenManager(t *testing.T) *gateway.AuthTokenManager {
	t.Helper()
	signingKey := []byte("test-signing-key-that-is-long-enough-for-hmac")
	return gateway.NewAuthTokenManager(signingKey, jwt.SigningMethodHS256, 60*time.Minute, 30*time.Minute)
}

func Test_AuthTokenManager_CreateToken_shouldReturnToken_whenValidInput(t *testing.T) {
	t.Parallel()

	// given
	m := newTestAuthTokenManager(t)

	// when
	token, err := m.CreateToken("user1", 1)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, token)
}

func Test_AuthTokenManager_ParseToken_shouldReturnUserInfo_whenValidToken(t *testing.T) {
	t.Parallel()

	// given
	m := newTestAuthTokenManager(t)
	token, err := m.CreateToken("user1", 1)
	require.NoError(t, err)

	// when
	userInfo, err := m.ParseToken(token)

	// then
	require.NoError(t, err)
	require.NotNil(t, userInfo)
	assert.Equal(t, 1, userInfo.UserID)
	assert.Equal(t, "user1", userInfo.LoginID)
}

func Test_AuthTokenManager_ParseToken_shouldReturnError_whenTokenIsInvalid(t *testing.T) {
	t.Parallel()

	// given
	m := newTestAuthTokenManager(t)

	// when
	userInfo, err := m.ParseToken("invalid-token-string")

	// then
	require.Error(t, err)
	assert.Nil(t, userInfo)
	assert.Contains(t, err.Error(), "parse token")
}

func Test_AuthTokenManager_ParseToken_shouldReturnError_whenTokenIsSignedWithDifferentKey(t *testing.T) {
	t.Parallel()

	// given
	creator := gateway.NewAuthTokenManager([]byte("original-key-that-is-long-enough"), jwt.SigningMethodHS256, 60*time.Minute, 30*time.Minute)
	parser := gateway.NewAuthTokenManager([]byte("different-key-that-is-long-enough"), jwt.SigningMethodHS256, 60*time.Minute, 30*time.Minute)
	token, err := creator.CreateToken("user1", 1)
	require.NoError(t, err)

	// when
	userInfo, err := parser.ParseToken(token)

	// then
	require.Error(t, err)
	assert.Nil(t, userInfo)
	assert.Contains(t, err.Error(), "parse token")
}

func Test_AuthTokenManager_ParseToken_shouldReturnError_whenTokenIsExpired(t *testing.T) {
	t.Parallel()

	// given
	m := gateway.NewAuthTokenManager([]byte("test-signing-key-that-is-long-enough-for-hmac"), jwt.SigningMethodHS256, -1*time.Minute, 30*time.Minute)
	token, err := m.CreateToken("user1", 1)
	require.NoError(t, err)

	// when
	userInfo, err := m.ParseToken(token)

	// then
	require.Error(t, err)
	assert.Nil(t, userInfo)
	assert.Contains(t, err.Error(), "parse token")
}

func Test_AuthTokenManager_ParseToken_shouldReturnExpiresAt(t *testing.T) {
	t.Parallel()

	// given
	m := newTestAuthTokenManager(t)
	token, err := m.CreateToken("user1", 1)
	require.NoError(t, err)

	// when
	userInfo, err := m.ParseToken(token)

	// then
	require.NoError(t, err)
	assert.False(t, userInfo.ExpiresAt.IsZero())
	assert.True(t, userInfo.ExpiresAt.After(time.Now()))
}

func Test_AuthTokenManager_RefreshToken_shouldReturnNewToken_whenRemainingTimeIsBelowThreshold(t *testing.T) {
	t.Parallel()

	// given
	signingKey := []byte("test-signing-key-that-is-long-enough-for-hmac")
	m := gateway.NewAuthTokenManager(signingKey, jwt.SigningMethodHS256, 2*time.Minute, 3*time.Minute)
	// remaining(~2min) < threshold(3min) → should refresh
	expiresAt := time.Now().Add(2 * time.Minute)

	// when
	newToken, err := m.RefreshToken("user1", 1, expiresAt)

	// then
	require.NoError(t, err)
	assert.NotEmpty(t, newToken)
}

func Test_AuthTokenManager_RefreshToken_shouldReturnEmpty_whenRemainingTimeIsAboveThreshold(t *testing.T) {
	t.Parallel()

	// given
	m := newTestAuthTokenManager(t)
	// remaining(~60min) > threshold(30min) → no refresh
	expiresAt := time.Now().Add(60 * time.Minute)

	// when
	newToken, err := m.RefreshToken("user1", 1, expiresAt)

	// then
	require.NoError(t, err)
	assert.Empty(t, newToken)
}
