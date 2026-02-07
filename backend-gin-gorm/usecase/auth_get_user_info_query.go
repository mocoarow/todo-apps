package usecase

import (
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

type AuthTokenParser interface {
	ParseToken(tokenString string) (*domain.UserInfo, error)
}

type AuthGetUserInfoQuery struct {
	authTokenParser AuthTokenParser
}

func NewAuthGetUserInfoQuery(authTokenParser AuthTokenParser) *AuthGetUserInfoQuery {
	return &AuthGetUserInfoQuery{
		authTokenParser: authTokenParser,
	}
}

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
