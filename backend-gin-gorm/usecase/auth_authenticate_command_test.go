package usecase_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
	"github.com/mocoarow/todo-apps/backend-gin-gorm/usecase"
)

func Test_AuthenticateCommand_Execute_shouldReturnToken_whenValidCredentials(t *testing.T) {
	t.Parallel()

	// given
	mockCreator := NewMockAuthTokenCreator(t)
	mockCreator.EXPECT().CreateToken("user1", 1).Return("access-token-123", nil).Once()
	cmd := usecase.NewAuthenticateCommand(mockCreator)
	input, err := domain.NewAuthenticateInput("user1", "password1")
	require.NoError(t, err)

	// when
	output, err := cmd.Execute(input)

	// then
	require.NoError(t, err)
	require.NotNil(t, output)
	assert.Equal(t, "access-token-123", output.AccessToken)
}

func Test_AuthenticateCommand_Execute_shouldReturnError_whenLoginIDFormatIsInvalid(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		loginID  string
		password string
	}{
		{
			name:     "loginID does not match pattern",
			loginID:  "admin",
			password: "password1",
		},
		{
			name:     "loginID is empty prefix",
			loginID:  "usr1",
			password: "password1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// given
			mockCreator := NewMockAuthTokenCreator(t)
			cmd := usecase.NewAuthenticateCommand(mockCreator)
			input := &domain.AuthenticateInput{LoginID: tt.loginID, Password: tt.password}

			// when
			output, err := cmd.Execute(input)

			// then
			require.Error(t, err)
			assert.Nil(t, output)
			assert.Contains(t, err.Error(), "authenticate user")
		})
	}
}

func Test_AuthenticateCommand_Execute_shouldReturnError_whenPasswordFormatIsInvalid(t *testing.T) {
	t.Parallel()

	// given
	mockCreator := NewMockAuthTokenCreator(t)
	cmd := usecase.NewAuthenticateCommand(mockCreator)
	input := &domain.AuthenticateInput{LoginID: "user1", Password: "wrongformat"}

	// when
	output, err := cmd.Execute(input)

	// then
	require.Error(t, err)
	assert.Nil(t, output)
	assert.Contains(t, err.Error(), "authenticate user")
}

func Test_AuthenticateCommand_Execute_shouldReturnError_whenCredentialsMismatch(t *testing.T) {
	t.Parallel()

	// given
	mockCreator := NewMockAuthTokenCreator(t)
	cmd := usecase.NewAuthenticateCommand(mockCreator)
	input := &domain.AuthenticateInput{LoginID: "user1", Password: "password2"}

	// when
	output, err := cmd.Execute(input)

	// then
	require.Error(t, err)
	assert.Nil(t, output)
	assert.Contains(t, err.Error(), "authenticate user")
}

func Test_AuthenticateCommand_Execute_shouldReturnError_whenCreateTokenFails(t *testing.T) {
	t.Parallel()

	// given
	mockCreator := NewMockAuthTokenCreator(t)
	mockCreator.EXPECT().CreateToken("user1", 1).Return("", fmt.Errorf("token creation failed")).Once()
	cmd := usecase.NewAuthenticateCommand(mockCreator)
	input, err := domain.NewAuthenticateInput("user1", "password1")
	require.NoError(t, err)

	// when
	output, err := cmd.Execute(input)

	// then
	require.Error(t, err)
	assert.Nil(t, output)
	assert.Contains(t, err.Error(), "create JWT")
}
