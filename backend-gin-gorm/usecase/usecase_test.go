package usecase_test

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/gateway"
)

var dbc *gateway.DBConnection

func TestMain(m *testing.M) {
	_ = godotenv.Load("../.env.test")

	tmpdbc, err := openTestMySQL()
	if err != nil {
		slog.Error("failed to open test MySQL", slog.Any("error", err))
		os.Exit(1)
	}

	dbc = tmpdbc

	code := m.Run()

	os.Exit(code)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func openTestMySQL() (*gateway.DBConnection, error) {
	host := getEnv("TEST_MYSQL_HOST", "127.0.0.1")
	portStr := getEnv("TEST_MYSQL_PORT", "3307")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid TEST_MYSQL_PORT: %w", err)
	}
	user := getEnv("TEST_MYSQL_USERNAME", "username")
	password := getEnv("TEST_MYSQL_PASSWORD", "password")
	database := getEnv("TEST_MYSQL_DATABASE", "test")

	db, err := gateway.OpenMySQL(&gateway.MySQLConfig{
		Username: user,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}, slog.LevelInfo, "test")
	if err != nil {
		return nil, fmt.Errorf("open test MySQL: %w", err)
	}

	return &gateway.DBConnection{DriverName: "mysql", DB: db}, nil
}

func cleanupTodoTable(t *testing.T, userID int) {
	t.Helper()
	if err := dbc.DB.Exec("DELETE FROM todo WHERE user_id = ?", userID).Error; err != nil {
		t.Fatalf("Failed to delete from table todo: %v", err)
	}
}
