package usecase

import (
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// AuthTokenManager combines token creation, parsing and refresh capabilities.
type AuthTokenManager interface {
	AuthTokenCreator
	AuthTokenParser
	AuthTokenRefresher
}

// AuthUsecase orchestrates authentication-related use cases.
type AuthUsecase struct {
	authenticateCommand *AuthenticateCommand
	getUserInfoQuery    *AuthGetUserInfoQuery
	refreshTokenQuery   *AuthRefreshTokenQuery
}

// NewAuthUsecase returns a new AuthUsecase with the given token manager.
func NewAuthUsecase(authTokenManager AuthTokenManager) *AuthUsecase {
	authenticateCommand := NewAuthenticateCommand(authTokenManager)
	getUserInfoQuery := NewAuthGetUserInfoQuery(authTokenManager)
	refreshTokenQuery := NewAuthRefreshTokenQuery(authTokenManager)
	return &AuthUsecase{
		authenticateCommand: authenticateCommand,
		getUserInfoQuery:    getUserInfoQuery,
		refreshTokenQuery:   refreshTokenQuery,
	}
}

// Authenticate validates credentials and returns a JWT access token.
func (u *AuthUsecase) Authenticate(input *domain.AuthenticateInput) (*domain.AuthenticateOutput, error) {
	output, err := u.authenticateCommand.Execute(input)
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}
	return output, nil
}

// GetUserInfo extracts user information from a JWT token.
func (u *AuthUsecase) GetUserInfo(input *domain.GetUserInfoInput) (*domain.GetUserInfoOutput, error) {
	output, err := u.getUserInfoQuery.Execute(input)
	if err != nil {
		return nil, fmt.Errorf("get user info: %w", err)
	}

	return output, nil
}

// RefreshToken checks if the token needs refresh and returns a new token if so.
func (u *AuthUsecase) RefreshToken(input *domain.RefreshTokenInput) (*domain.RefreshTokenOutput, error) {
	output, err := u.refreshTokenQuery.Execute(input)
	if err != nil {
		return nil, fmt.Errorf("refresh token: %w", err)
	}

	return output, nil
}
