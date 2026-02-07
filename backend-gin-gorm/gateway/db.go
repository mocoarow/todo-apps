// Package gateway provides infrastructure adapters for external systems
// such as databases, JWT authentication, logging/telemetry, and OS signal handling.
package gateway

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"gorm.io/gorm"
)

// DBConfig holds the database driver name and driver-specific configuration.
type DBConfig struct {
	DriverName string       `yaml:"driverName"`
	MySQL      *MySQLConfig `yaml:"mysql"`
}

// DBConnection wraps an active GORM database connection along with its dialect.
type DBConnection struct {
	DriverName string
	Dialect    DialectRDBMS
	DB         *gorm.DB
}

// InitDBFunc is a function type that initializes a database connection for a specific driver.
type InitDBFunc func(context.Context, *DBConfig, slog.Level, string) (DialectRDBMS, *gorm.DB, *sql.DB, error)

// InitDB initializes a database connection based on the configured driver.
// It returns the connection and a cleanup function to close the underlying sql.DB.
func InitDB(ctx context.Context, dbConfig *DBConfig, logConfig *LogConfig, appName string) (*DBConnection, func(), error) {
	initDBs := map[string]InitDBFunc{
		"mysql": initDBMySQL,
	}

	initDBFunc, ok := initDBs[dbConfig.DriverName]
	if !ok {
		return nil, nil, fmt.Errorf("invalid database driver: %s", dbConfig.DriverName)
	}
	dbLogLevel := slog.LevelWarn
	if level, ok := logConfig.Levels["db"]; ok {
		dbLogLevel = stringToLogLevel(level)
	}

	dialect, db, sqlDB, err := initDBFunc(ctx, dbConfig, dbLogLevel, appName)
	if err != nil {
		return nil, nil, fmt.Errorf("init DB: %w", err)
	}

	dbConn := DBConnection{
		DriverName: dbConfig.DriverName,
		Dialect:    dialect,
		DB:         db,
	}

	return &dbConn, func() {
		if err := sqlDB.Close(); err != nil {
			slog.Error("failed to close sqlDB", "error", err)
		}
	}, nil
}

// DialectRDBMS abstracts database dialect differences (e.g. default values).
type DialectRDBMS interface {
	Name() string
	BoolDefaultValue() string
}
