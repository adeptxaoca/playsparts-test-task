package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/tern/migrate"

	"part_handler/internal/app/config"
	"part_handler/internal/pkg/errors"
)

type DB struct {
	Pool *pgxpool.Pool
}

// Database configuration and connection creation
func Setup(ctx context.Context, conf *config.Config) (*DB, int32, error) {
	db, err := connect(ctx, conf)
	if err != nil {
		return nil, 0, errors.Wrap(err, "Unable to connect to database")
	}

	ver, err := db.migrateDatabase(ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "Unable to use migrations")
	}

	return db, ver, nil
}

// Connect creates a new Pool and immediately establishes one connection
func connect(ctx context.Context, conf *config.Config) (*DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s",
		conf.Database.User, conf.Database.Pass, conf.Database.Addr, conf.Database.Name)
	pool, err := pgxpool.Connect(ctx, connString)
	return &DB{Pool: pool}, err
}

// Run migrations to the database
func (db *DB) migrateDatabase(ctx context.Context) (int32, error) {
	conn, err := db.acquire(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Release()

	migrator, err := migrate.NewMigrator(ctx, conn.Conn(), "public.schema_version")
	if err != nil {
		return 0, errors.InternalError.Wrap(err, "Unable to create a migrator")
	}

	err = migrator.LoadMigrations("./internal/app/migrations/")
	if err != nil {
		return 0, errors.InternalError.Wrap(err, "Unable to load migrations")
	}

	err = migrator.Migrate(ctx)
	if err != nil {
		return 0, errors.InternalError.Wrap(err, "Unable to migrate")
	}

	ver, err := migrator.GetCurrentVersion(ctx)
	if err != nil {
		return 0, errors.InternalError.Wrap(err, "Unable to get current schema version")
	}

	return ver, nil
}

func (db *DB) acquire(ctx context.Context) (*pgxpool.Conn, error) {
	conn, err := db.Pool.Acquire(ctx)
	if err != nil {
		return nil, errors.InternalError.Wrap(err, "Unable to acquire a database connection")
	}
	return conn, nil
}

func pgError(err error) error {
	if err == nil {
		return nil
	}

	origErr := errors.Cause(err)
	if origErr == pgx.ErrNoRows {
		return errors.NotFound.New("record was not found")
	}

	// perhaps best use pgErr.ConstraintName
	if pgErr, ok := origErr.(*pgconn.PgError); ok {
		switch pgErr.Code {
		case "23000":
			return errors.IntegrityRestrictionError.Newf("integrity constraint violation (%s)", pgErr.Detail)
		case "23001":
			return errors.IntegrityRestrictionError.Newf("restrict violation (%s)", pgErr.Detail)
		case "23502":
			return errors.IntegrityRestrictionError.Newf("not null violation (%s)", pgErr.Detail)
		case "23503":
			return errors.IntegrityRestrictionError.Newf("foreign key violation (%s)", pgErr.Detail)
		case "23505":
			return errors.IntegrityRestrictionError.Newf("unique violation (%s)", pgErr.Detail)
		case "23514":
			return errors.IntegrityRestrictionError.Newf("check violation (%s)", pgErr.Detail)
		case "23P01":
			return errors.IntegrityRestrictionError.Newf("exclusion violation (%s)", pgErr.Detail)
		default:
			return errors.DatabaseError.Wrap(err, "database error")
		}
	}

	return errors.InternalError.Wrap(err, "undefined pg error")
}
