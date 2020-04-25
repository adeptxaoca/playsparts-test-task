package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/tern/migrate"

	"part_handler/internal/app/config"
	"part_handler/internal/pkg/errors"
)

type DB struct {
	Pool *pgxpool.Pool
}

// Database configuration and connection creation
func Setup(ctx context.Context, conf *config.Config) (*DB, error) {
	db, err := connect(ctx, conf)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to database")
	}

	if err := db.migrateDatabase(context.Background()); err != nil {
		return nil, errors.Wrap(err, "Unable to use migrations")
	}

	return db, nil
}

// Connect creates a new Pool and immediately establishes one connection
func connect(ctx context.Context, conf *config.Config) (*DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s",
		conf.Database.User, conf.Database.Pass, conf.Database.Addr, conf.Database.Name)
	pool, err := pgxpool.Connect(ctx, connString)
	return &DB{Pool: pool}, err
}

// Run migrations to the database
func (db *DB) migrateDatabase(ctx context.Context) error {
	conn, err := db.acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	migrator, err := migrate.NewMigrator(ctx, conn.Conn(), "public.schema_version")
	if err != nil {
		return errors.InternalError.Wrap(err, "Unable to create a migrator")
	}

	err = migrator.LoadMigrations("./internal/app/database/migrations/")
	if err != nil {
		return errors.InternalError.Wrap(err, "Unable to load migrations")
	}

	err = migrator.Migrate(ctx)
	if err != nil {
		return errors.InternalError.Wrap(err, "Unable to migrate")
	}

	ver, err := migrator.GetCurrentVersion(ctx)
	if err != nil {
		return errors.InternalError.Wrap(err, "Unable to get current schema version")
	}

	log.Printf("Migration done. Current schema version: %v\n", ver)

	return nil
}

func (db *DB) acquire(ctx context.Context) (*pgxpool.Conn, error) {
	conn, err := db.Pool.Acquire(ctx)
	if err != nil {
		return nil, errors.InternalError.Wrap(err, "Unable to acquire a database connection")
	}
	return conn, nil
}
