package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	ErrDBNotInitialized = fmt.Errorf("database is not initialized")
)

type Postgres struct {
	DB *sqlx.DB
}

const (
	connectAttempts = 15
)

// New creates a new Postgres instance
func New(cfg *DbConfig) (Postgres, error) {
	var (
		db        *sql.DB
		migration *migrate.Migrate
		driver    database.Driver
		err       error
	)

	for attempt := 0; attempt < connectAttempts; attempt++ {
		time.Sleep(time.Second)

		db, err = sql.Open("postgres", cfg.ConnectionString)
		if err != nil {
			continue
		}
	}

	driver, err = postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return Postgres{}, fmt.Errorf("failed to create postgres driver: %w", err)
	}

	fs, err := os.ReadDir(cfg.MigrationsPath)
	if err != nil {
		return Postgres{}, fmt.Errorf("failed to stat migrations path: %w", err)
	}

	if len(fs) == 0 {
		return Postgres{}, fmt.Errorf("no migrations found in %s", cfg.MigrationsPath)
	}

	for _, f := range fs {
		if !f.IsDir() {
			continue
		}

		fmt.Println(f.Name())
	}

	migrationsPath := fmt.Sprintf("file://%s", cfg.MigrationsPath)
	migration, err = migrate.NewWithDatabaseInstance(
		migrationsPath,
		"postgres", driver)
	if err != nil {
		return Postgres{}, fmt.Errorf("failed to create migrations instance: %w", err)
	}

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return Postgres{}, fmt.Errorf("failed to apply migrations: %w", err)
	}

	sdb := sqlx.NewDb(db, "postgres")
	sdb.SetMaxOpenConns(cfg.MaxOpenConns)

	err = sdb.Ping()
	if err != nil {
		return Postgres{}, fmt.Errorf("failed to ping database: %w", err)
	}

	return Postgres{
		DB: sdb,
	}, nil
}

func (p *Postgres) Close() error {
	if p.DB == nil {
		return ErrDBNotInitialized
	}

	return p.DB.Close()
}

func (p Postgres) CheckHealth(ctx context.Context) error {
	if p.DB == nil {
		return ErrDBNotInitialized
	}

	return p.DB.PingContext(ctx)
}

func (p Postgres) CheckReadiness(ctx context.Context) error {
	if p.DB == nil {
		return ErrDBNotInitialized
	}

	return p.DB.PingContext(ctx)
}
