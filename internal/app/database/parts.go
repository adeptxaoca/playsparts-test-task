package database

import (
	"context"

	"part_handler/internal/app/models/part"
)

// Call the "Create" function of the model "part"
func (db *DB) CreatePart(ctx context.Context, in *part.Part) (*part.Part, error) {
	conn, err := db.acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	out, err := part.Create(ctx, conn, in)
	return out, pgError(err)
}

// Call the "Read" function of the model "part"
func (db *DB) ReadPart(ctx context.Context, id uint64) (*part.Part, error) {
	conn, err := db.acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	out, err := part.Read(ctx, conn, id)
	return out, pgError(err)
}

// Call the "Update" function of the model "part"
func (db *DB) UpdatePart(ctx context.Context, in *part.Part) (*part.Part, error) {
	conn, err := db.acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	out, err := part.Update(ctx, conn, in)
	return out, pgError(err)
}

// Call the "Delete" function of the model "part"
func (db *DB) DeletePart(ctx context.Context, id uint64) error {
	conn, err := db.acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	err = part.Delete(ctx, conn, id)
	return pgError(err)
}
