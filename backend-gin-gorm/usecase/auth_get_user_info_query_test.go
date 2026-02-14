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

func Test_AuthGetUserInfoQuery_Execute_shouldReturnUserInfo_whenValidToken(t *testing.T) {
	t.Parallel()

	// given
	mockParser := NewMockAuthTokenParser(t)
	userInfo, err := domain.NewUserInfo(1, "user1", time.Now().Add(60*time.Minute))
	require.NoError(t, err)
	mockParser.EXPECT().ParseToken("valid-token").Return(userInfo, nil).Once()
	query := usecase.NewAuthGetUserInfoQuery(mockParser)
	input, err := domain.NewGetUserInfoInput("valid-token")
	require.NoError(t, err)

	// when
	output, err := query.Execute(input)

	// then
	require.NoError(t, err)
	require.NotNil(t, output)
	assert.Equal(t, 1, output.UserInfo.UserID)
	assert.Equal(t, "user1", output.UserInfo.LoginID)
}

func Test_AuthGetUserInfoQuery_Execute_shouldReturnError_whenParseTokenFails(t *testing.T) {
	t.Parallel()

	// given
	mockParser := NewMockAuthTokenParser(t)
	mockParser.EXPECT().ParseToken("invalid-token").Return(nil, errors.New("token parse failed")).Once()
	query := usecase.NewAuthGetUserInfoQuery(mockParser)
	input, err := domain.NewGetUserInfoInput("invalid-token")
	require.NoError(t, err)

	// when
	output, err := query.Execute(input)

	// then
	require.Error(t, err)
	assert.Nil(t, output)
	assert.Contains(t, err.Error(), "parse token")
}
