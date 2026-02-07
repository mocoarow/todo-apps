package usecase

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

type AuthTokenCreator interface {
	CreateToken(loginID string, userID int) (string, error)
}

type AuthenticateCommand struct {
	authTokenCreator AuthTokenCreator
	regexpUserID     *regexp.Regexp
	regexpPassword   *regexp.Regexp
}

func NewAuthenticateCommand(authTokenCreator AuthTokenCreator) *AuthenticateCommand {
	return &AuthenticateCommand{
		authTokenCreator: authTokenCreator,
		regexpUserID:     regexp.MustCompile(`^user([\d]+)$`),
		regexpPassword:   regexp.MustCompile(`^password([\d]+)$`),
	}
}

func (c *AuthenticateCommand) Execute(input *domain.AuthenticateInput) (*domain.AuthenticateOutput, error) {
	userID, err := c.authenticate(input.LoginID, input.Password)
	if err != nil {
		return nil, fmt.Errorf("authenticate user: %w: %w", domain.ErrUnauthenticated, err)
	}

	accessToken, err := c.authTokenCreator.CreateToken(input.LoginID, userID)
	if err != nil {
		return nil, fmt.Errorf("create JWT: %w", err)
	}

	output, err := domain.NewAuthenticateOutput(accessToken)
	if err != nil {
		return nil, fmt.Errorf("create authenticate output: %w", err)
	}

	return output, nil
}

func (c *AuthenticateCommand) authenticate(loginID string, password string) (int, error) {
	userIDMatches := c.regexpUserID.FindStringSubmatch(loginID)
	if len(userIDMatches) != 2 {
		return 0, fmt.Errorf("invalid login ID format")
	}
	userIDStr := userIDMatches[1]

	passwordMatches := c.regexpPassword.FindStringSubmatch(password)
	if len(passwordMatches) != 2 {
		return 0, fmt.Errorf("invalid password format")
	}
	passwordNum := passwordMatches[1]

	if userIDStr != passwordNum {
		return 0, fmt.Errorf("invalid login ID or password")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, fmt.Errorf("convert userID to int: %w", err)
	}

	return userID, nil
}
