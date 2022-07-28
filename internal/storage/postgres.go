package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	// this import is required to apply PostgreSQL migrations.
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	// this import is required to load migrations from *.sql files.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"

	// this import is required to open connection with the PostgreSQL database.
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/chiefcake/apod/internal/config"
)

// Postgres contains all methods for querying data from the PostgreSQL database.
type Postgres struct {
	*sqlx.DB
}

// NewPostgres connects to the PostgreSQL database, applies database migrations and returns an opened connection or an error.
func NewPostgres(ctx context.Context, cfg *config.Config) (*Postgres, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", cfg.PostgresUsername,
		cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresDatabase, cfg.PostgresSSLMode)

	log.Println("Connecting to PostgreSQL database...")

	db, err := sqlx.ConnectContext(ctx, "postgres", url)
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to the PostgreSQL")
	}

	migrator, err := migrate.New(cfg.MigrationsDirectory, url)
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize migrator")
	}

	log.Println("Applying database migrations...")

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, errors.Wrap(err, "could not apply migrations")
	}

	return &Postgres{
		DB: db,
	}, nil
}
