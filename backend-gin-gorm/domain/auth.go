package domain

import "fmt"

type UserInfo struct {
	UserID  int    `validate:"required,gt=0"`
	LoginID string `validate:"required"`
}

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

type GetUserInfoInput struct {
	TokenString string `validate:"required"`
}

func NewGetUserInfoInput(tokenString string) (*GetUserInfoInput, error) {
	m := &GetUserInfoInput{
		TokenString: tokenString,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate get user info input: %w", err)
	}
	return m, nil
}

type GetUserInfoOutput struct {
	UserInfo *UserInfo `validate:"required"`
}

func NewGetUserInfoOutput(userInfo *UserInfo) (*GetUserInfoOutput, error) {
	m := &GetUserInfoOutput{
		UserInfo: userInfo,
	}
	if err := ValidateStruct(m); err != nil {
		return nil, fmt.Errorf("validate get user info output: %w", err)
	}
	return m, nil
}
