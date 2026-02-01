package gateway_test

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"testing"

	"gorm.io/gorm"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/gateway"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	tmpdb, err := OpenTestMySQL()
	if err != nil {
		slog.Error("failed to open test MySQL", slog.String("error", err.Error()))
		os.Exit(1)
	}

	db = tmpdb

	// run tests
	code := m.Run()

	os.Exit(code)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// SetupTestDB sets up a test database connection and runs migrations
func OpenTestMySQL() (*gorm.DB, error) {
	host := getEnv("TEST_MYSQL_HOST", "127.0.0.1")
	portStr := getEnv("TEST_MYSQL_PORT", "3307")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid TEST_MYSQL_PORT: %w", err)
	}
	user := getEnv("TEST_MYSQL_USERNAME", "username")
	password := getEnv("TEST_MYSQL_PASSWORD", "password")
	database := getEnv("TEST_MYSQL_DATABASE", "test")
	logLevelStr := getEnv("TEST_LOG_LEVEL", "INFO")
	logLevel := slog.LevelInfo
	if logLevelStr == "DEBUG" {
		logLevel = slog.LevelDebug
	}

	db, err := gateway.OpenMySQL(&gateway.MySQLConfig{
		Username: user,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}, logLevel, "test")
	if err != nil {
		return nil, fmt.Errorf("open test MySQL: %w", err)
	}

	return db, nil
}
