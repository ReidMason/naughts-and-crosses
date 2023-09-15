package migrations

import (
	"database/sql"
	"log"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func getDatabaseMigrationInstance(db *sql.DB) *migrate.Migrate {
	driverName := "postgres"
	migrationsFilepath := "file://internal/migrations/sqlMigrations"

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsFilepath,
		driverName, driver)
	if err != nil {
		log.Panic(err)
	}

	return m
}

func Migrate(db *sql.DB) {
	migrationInstance := getDatabaseMigrationInstance(db)
	slog.Info("Performing database migration")
	err := migrationInstance.Up()
	if err != nil && err.Error() != migrate.ErrNoChange.Error() {
		slog.Error("Failed to migrate", "error", err)
		log.Panic()
	} else if err.Error() == migrate.ErrNoChange.Error() {
		slog.Info("Database migration complete - No changes")
	} else {
		slog.Info("Database migration complete")
	}
}
