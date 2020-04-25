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

	return part.Create(ctx, conn, in)
}

// Call the "Read" function of the model "part"
func (db *DB) ReadPart(ctx context.Context, id uint64) (*part.Part, error) {
	conn, err := db.acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	return part.Read(ctx, conn, id)
}

// Call the "Update" function of the model "part"
func (db *DB) UpdatePart(ctx context.Context, in *part.Part) (*part.Part, error) {
	conn, err := db.acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	return part.Update(ctx, conn, in)
}

// Call the "Delete" function of the model "part"
func (db *DB) DeletePart(ctx context.Context, id uint64) error {
	conn, err := db.acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	return part.Delete(ctx, conn, id)
}
