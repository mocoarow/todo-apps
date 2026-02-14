package usecase_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/usecase"
)

func Test_AuthRefreshTokenQuery_Execute_shouldReturnNewToken_whenRefreshNeeded(t *testing.T) {
	t.Parallel()

	// given
	expiresAt := time.Now().Add(5 * time.Minute)
	mockRefresher := NewMockAuthTokenRefresher(t)
	mockRefresher.EXPECT().RefreshToken("user1", 1, expiresAt).Return("new-token", nil).Once()
	query := usecase.NewAuthRefreshTokenQuery(mockRefresher)
	input, err := domain.NewRefreshTokenInput("user1", 1, expiresAt)
	require.NoError(t, err)

	// when
	output, err := query.Execute(input)

	// then
	require.NoError(t, err)
	require.NotNil(t, output)
	assert.Equal(t, "new-token", output.NewAccessToken)
}

func Test_AuthRefreshTokenQuery_Execute_shouldReturnEmpty_whenNoRefreshNeeded(t *testing.T) {
	t.Parallel()

	// given
	expiresAt := time.Now().Add(60 * time.Minute)
	mockRefresher := NewMockAuthTokenRefresher(t)
	mockRefresher.EXPECT().RefreshToken("user1", 1, expiresAt).Return("", nil).Once()
	query := usecase.NewAuthRefreshTokenQuery(mockRefresher)
	input, err := domain.NewRefreshTokenInput("user1", 1, expiresAt)
	require.NoError(t, err)

	// when
	output, err := query.Execute(input)

	// then
	require.NoError(t, err)
	require.NotNil(t, output)
	assert.Empty(t, output.NewAccessToken)
}

func Test_AuthRefreshTokenQuery_Execute_shouldReturnError_whenRefreshFails(t *testing.T) {
	t.Parallel()

	// given
	expiresAt := time.Now().Add(5 * time.Minute)
	mockRefresher := NewMockAuthTokenRefresher(t)
	mockRefresher.EXPECT().RefreshToken("user1", 1, expiresAt).Return("", errors.New("token refresh failed")).Once()
	query := usecase.NewAuthRefreshTokenQuery(mockRefresher)
	input, err := domain.NewRefreshTokenInput("user1", 1, expiresAt)
	require.NoError(t, err)

	// when
	output, err := query.Execute(input)

	// then
	require.Error(t, err)
	assert.Nil(t, output)
	assert.Contains(t, err.Error(), "refresh token")
}
