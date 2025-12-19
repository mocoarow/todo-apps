package domain

import (
	"github.com/go-playground/validator/v10"
)

const (
	LoggerNameKey = "logger_name"
	AppName       = "backend-gin-gorm"
)

var (
	Validator = validator.New()
)
