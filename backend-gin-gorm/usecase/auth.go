package usecase

import (
	"fmt"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

type AuthTokenManager interface {
	AuthTokenCreator
	AuthTokenParser
}

type AuthUsecase struct {
	authenticateCommand *AuthenticateCommand
	getUserInfoQuery    *AuthGetUserInfoQuery
}

func NewAuthUsecase(authTokenManager AuthTokenManager) *AuthUsecase {
	authenticateCommand := NewAuthenticateCommand(authTokenManager)
	getUserInfoQuery := NewAuthGetUserInfoQuery(authTokenManager)
	return &AuthUsecase{
		authenticateCommand: authenticateCommand,
		getUserInfoQuery:    getUserInfoQuery,
	}
}

func (u *AuthUsecase) Authenticate(input *domain.AuthenticateInput) (*domain.AuthenticateOutput, error) {
	output, err := u.authenticateCommand.Execute(input)
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}
	return output, nil
}

func (u *AuthUsecase) GetUserInfo(input *domain.GetUserInfoInput) (*domain.GetUserInfoOutput, error) {
	output, err := u.getUserInfoQuery.Execute(input)
	if err != nil {
		return nil, fmt.Errorf("get user info: %w", err)
	}

	return output, nil
}
