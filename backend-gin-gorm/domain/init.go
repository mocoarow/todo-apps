package domain

import "github.com/go-playground/validator/v10"

const (
	// LoggerNameKey is the structured log key used to identify the logger name.
	LoggerNameKey = "logger_name"
	// AppName is the application name used for logging and tracing.
	AppName = "backend-gin-gorm"
)

var (
	v = validator.New()
)

// ValidateStruct validates the given struct using the go-playground/validator tags.
func ValidateStruct(s interface{}) error {
	return v.Struct(s) //nolint:wrapcheck
}
