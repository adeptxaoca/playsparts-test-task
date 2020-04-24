package database

import (
	"context"

	"part_handler/internal/part_handler/models/part"
	"part_handler/pkg/errors"
)

// Insert a new part in the table "parts"
func (db *DB) CreatePart(ctx context.Context, p *part.Part) error {
	conn, err := db.acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	err = conn.QueryRow(ctx, "INSERT INTO parts (manufacturer_id, name, vendor_code) VALUES($1, $2, "+
		"$3) RETURNING id, created_at", p.ManufacturerId, p.Name, p.VendorCode).Scan(&p.Id, &p.CreatedAt)

	return err
}

// Select a part by id
func (db *DB) ReadPart(ctx context.Context, id uint64) (*part.Part, error) {
	conn, err := db.acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	var p part.Part
	err = conn.QueryRow(ctx, "SELECT * FROM parts WHERE id = $1", id).
		Scan(&p.Id, &p.ManufacturerId, &p.Name, &p.VendorCode, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)

	return &p, err
}

// Update a part by id
func (db *DB) UpdatePart(ctx context.Context, p *part.Part) (*part.Part, error) {
	conn, err := db.acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	err = conn.QueryRow(ctx, "UPDATE parts SET manufacturer_id = $1, name = $2, vendor_code = $3 "+
		"WHERE id = $4 RETURNING updated_at", p.ManufacturerId, p.Name, p.VendorCode, p.Id).Scan(&p.UpdatedAt)

	return p, err
}

// Delete a part by id
func (db *DB) DeletePart(ctx context.Context, id uint64) error {
	conn, err := db.acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	ct, err := conn.Exec(ctx, "DELETE FROM parts WHERE id = $1", id)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return errors.NotFound.Wrap(err, "part not found")
	}

	return nil
}
