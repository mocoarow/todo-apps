package usecase

import (
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// AuthTokenParser parses a JWT token string and extracts user information.
type AuthTokenParser interface {
	ParseToken(tokenString string) (*domain.UserInfo, error)
}

// AuthGetUserInfoQuery retrieves user info by parsing a JWT token.
type AuthGetUserInfoQuery struct {
	authTokenParser AuthTokenParser
}

// NewAuthGetUserInfoQuery returns a new AuthGetUserInfoQuery.
func NewAuthGetUserInfoQuery(authTokenParser AuthTokenParser) *AuthGetUserInfoQuery {
	return &AuthGetUserInfoQuery{
		authTokenParser: authTokenParser,
	}
}

// Execute parses the token from input and returns the associated user info.
func (u *AuthGetUserInfoQuery) Execute(input *domain.GetUserInfoInput) (*domain.GetUserInfoOutput, error) {
	userInfo, err := u.authTokenParser.ParseToken(input.TokenString)
	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}

	output, err := domain.NewGetUserInfoOutput(userInfo)
	if err != nil {
		return nil, fmt.Errorf("create get user info output: %w", err)
	}
	return output, nil
}
