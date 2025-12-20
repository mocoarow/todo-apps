package config

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"go.yaml.in/yaml/v4"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/domain"
)

type ServerConfig struct {
	HTTPPort             int `yaml:"httpPort" validate:"required"`
	MetricsPort          int `yaml:"metricsPort" validate:"required"`
	ReadHeaderTimeoutSec int `yaml:"readHeaderTimeoutSec" validate:"gte=1"`
}

type Config struct {
	Server *ServerConfig `yaml:"server" validate:"required"`
}

//go:embed config.yml
var config embed.FS

// ExpandEnvWithDefaults expands environment variables in the format VAR_NAME:-default_value.
func ExpandEnvWithDefaults(varName string) string {
	// Check if it contains :-
	if strings.Contains(varName, ":-") {
		parts := strings.SplitN(varName, ":-", 2)
		name := parts[0]
		defaultValue := parts[1]

		if value := os.Getenv(name); value != "" {
			return value
		}

		return defaultValue
	}

	// Simple variable expansion
	return os.Getenv(varName)
}

func LoadConfig() (*Config, error) {
	filename := "config.yml"
	confContent, err := config.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("config.ReadFile. filename: %s, err: %w", filename, err)
	}

	confContent = []byte(os.Expand(string(confContent), ExpandEnvWithDefaults))
	var conf Config
	if err := yaml.Unmarshal(confContent, &conf); err != nil {
		return nil, fmt.Errorf("yaml.Unmarshal. filename: %s, err: %w", filename, err)
	}

	if err := domain.Validator.Struct(&conf); err != nil {
		return nil, fmt.Errorf("Validator.Struct. filename: %s, err: %w", filename, err)
	}

	return &conf, nil
}
