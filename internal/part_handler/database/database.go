package database

import (
	"context"

	"github.com/jackc/pgx/v4"

	"part_handler/internal/part_handler/models/part"
)

type DB struct {
	Conn *pgx.Conn
}

// Connect establishes a connection with a PostgreSQL server with a connection string
func Connect(ctx context.Context, connString string) (*DB, error) {
	conn, err := pgx.Connect(ctx, connString)
	return &DB{Conn: conn}, err
}

// TODO: Insert a new part in the table "parts"
func (db *DB) CreatePart(ctx context.Context, part *part.Part) (*part.Part, error) {
	_, err := db.Conn.Exec(ctx, "INSERT INTO parts (manufacturer_id, name, vendor_code) VALUES($1, $2, $3)",
		part.ManufacturerId, part.Name, part.VendorCode)
	return part, err
}

// TODO: Select a part by id
func (db *DB) ReadPart(ctx context.Context, id uint64) (*part.Part, error) {
	var p part.Part
	err := db.Conn.QueryRow(ctx, "SELECT * FROM parts WHERE id = $1", id).
		Scan(&p.Id, &p.ManufacturerId, &p.Name, &p.VendorCode, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt)
	return &p, err
}
