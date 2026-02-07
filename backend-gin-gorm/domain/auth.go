package domain

import (
	"errors"
	"fmt"
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

// UserInfo represents an authenticated user's identity.
type UserInfo struct {
	UserID  int    `validate:"required,gt=0"`
	LoginID string `validate:"required"`
}

// NewUserInfo creates a validated UserInfo.
func NewUserInfo(userID int, loginID string) (*UserInfo, error) {
	m := &UserInfo{
		UserID:  userID,
		LoginID: loginID,
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
