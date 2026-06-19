package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"opspilot/backend/internal/config"
)

func main() {
	direction := flag.String("direction", "up", "migration direction: up or down")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := config.Load()

	if cfg.DatabaseURL == "" {
		logger.Error("DATABASE_URL is required")
		os.Exit(1)
	}

	migration, err := migrate.New("file://migrations", cfg.DatabaseURL)
	if err != nil {
		logger.Error("failed to initialize migrations", "error", err)
		os.Exit(1)
	}
	defer migration.Close()

	switch *direction {
	case "up":
		err = migration.Up()
	case "down":
		err = migration.Steps(-1)
	default:
		logger.Error("unsupported migration direction", "direction", *direction)
		os.Exit(1)
	}

	if err != nil && err != migrate.ErrNoChange {
		logger.Error("migration failed", "error", err)
		os.Exit(1)
	}

	if err == migrate.ErrNoChange {
		logger.Info("no migration changes")
		return
	}

	logger.Info("migration completed", "direction", *direction)
}
