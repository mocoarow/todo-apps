package domain

import (
	"errors"
	"fmt"
	"time"
)

// ErrUnauthenticated is returned when authentication fails due to invalid credentials.
var ErrUnauthenticated = errors.New("unauthenticated")

// AuthenticateInput holds the login credentials for authentication.
type AuthenticateInput struct {
	LoginID  string `validate:"required"`
	Password string `validate:"required"`
}

// NewAuthenticateInput creates a validated AuthenticateInput.
func NewAuthenticateInput(loginID string, password string) (*AuthenticateInput, error) {
	m := &AuthenticateInput{
		LoginID:  loginID,
		Password: password,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate authenticate input: %w", err)
	}
	return m, nil
}

// AuthenticateOutput holds the access token returned after successful authentication.
type AuthenticateOutput struct {
	AccessToken string `validate:"required"`
}

// NewAuthenticateOutput creates a validated AuthenticateOutput.
func NewAuthenticateOutput(accessToken string) (*AuthenticateOutput, error) {
	m := &AuthenticateOutput{
		AccessToken: accessToken,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate authenticate output: %w", err)
	}
	return m, nil
}

// UserInfo represents an authenticated user's identity extracted from a JWT token.
type UserInfo struct {
	UserID    int       `validate:"required,gt=0"`
	LoginID   string    `validate:"required"`
	ExpiresAt time.Time `validate:"required"`
}

// NewUserInfo creates a validated UserInfo.
func NewUserInfo(userID int, loginID string, expiresAt time.Time) (*UserInfo, error) {
	m := &UserInfo{
		UserID:    userID,
		LoginID:   loginID,
		ExpiresAt: expiresAt,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate user info: %w", err)
	}
	return m, nil
}

// GetUserInfoInput holds the JWT token string to be parsed.
type GetUserInfoInput struct {
	TokenString string `validate:"required"`
}

// NewGetUserInfoInput creates a validated GetUserInfoInput.
func NewGetUserInfoInput(tokenString string) (*GetUserInfoInput, error) {
	m := &GetUserInfoInput{
		TokenString: tokenString,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate get user info input: %w", err)
	}
	return m, nil
}

// GetUserInfoOutput holds the user info extracted from a JWT token.
type GetUserInfoOutput struct {
	UserInfo *UserInfo `validate:"required"`
}

// NewGetUserInfoOutput creates a validated GetUserInfoOutput.
func NewGetUserInfoOutput(userInfo *UserInfo) (*GetUserInfoOutput, error) {
	m := &GetUserInfoOutput{
		UserInfo: userInfo,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate get user info output: %w", err)
	}
	return m, nil
}

// RefreshTokenInput holds the parsed claims needed for a token refresh check.
type RefreshTokenInput struct {
	LoginID   string    `validate:"required"`
	UserID    int       `validate:"required,gt=0"`
	ExpiresAt time.Time `validate:"required"`
}

// NewRefreshTokenInput creates a validated RefreshTokenInput.
func NewRefreshTokenInput(loginID string, userID int, expiresAt time.Time) (*RefreshTokenInput, error) {
	m := &RefreshTokenInput{
		LoginID:   loginID,
		UserID:    userID,
		ExpiresAt: expiresAt,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate refresh token input: %w", err)
	}
	return m, nil
}

// RefreshTokenOutput holds the result of a token refresh attempt.
// NewAccessToken is empty if no refresh was needed.
type RefreshTokenOutput struct {
	NewAccessToken string
}

// NewRefreshTokenOutput creates a RefreshTokenOutput.
func NewRefreshTokenOutput(newAccessToken string) *RefreshTokenOutput {
	return &RefreshTokenOutput{
		NewAccessToken: newAccessToken,
	}
}
