package usecase

import (
	"fmt"
	"time"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

// AuthTokenRefresher checks if a token needs refresh and issues a new one if so.
type AuthTokenRefresher interface {
	RefreshToken(loginID string, userID int, expiresAt time.Time) (string, error)
}

// AuthRefreshTokenQuery handles token refresh logic.
type AuthRefreshTokenQuery struct {
	authTokenRefresher AuthTokenRefresher
}

// NewAuthRefreshTokenQuery returns a new AuthRefreshTokenQuery.
func NewAuthRefreshTokenQuery(authTokenRefresher AuthTokenRefresher) *AuthRefreshTokenQuery {
	return &AuthRefreshTokenQuery{
		authTokenRefresher: authTokenRefresher,
	}
}

// Execute checks if the token needs refresh and returns a new token if so.
func (q *AuthRefreshTokenQuery) Execute(input *domain.RefreshTokenInput) (*domain.RefreshTokenOutput, error) {
	newToken, err := q.authTokenRefresher.RefreshToken(input.LoginID, input.UserID, input.ExpiresAt)
	if err != nil {
		return nil, fmt.Errorf("refresh token: %w", err)
	}

	return domain.NewRefreshTokenOutput(newToken), nil
}
