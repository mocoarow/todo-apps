package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/mocoarow/todo-apps/backend-gin-gorm/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	if _, err := config.LoadConfig(); err != nil {
		return fmt.Errorf("LoadConfig: %w", err)
	}

	slog.Default().Info("Hello World")

	return nil
}
