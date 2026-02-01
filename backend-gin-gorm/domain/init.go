package domain

import "github.com/go-playground/validator/v10"

const (
	LoggerNameKey = "logger_name"
	AppName       = "backend-gin-gorm"
)

var (
	v = validator.New()
)

func ValidateStruct(s interface{}) error {
	return v.Struct(s) //nolint:wrapcheck
}
