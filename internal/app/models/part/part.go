package part

import (
	"context"

	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"

	pb "part_handler/internal/pkg/api/v1"
	"part_handler/internal/pkg/errors"
)

type Part struct {
	Id             uint64             `json:"id"`
	ManufacturerId uint64             `json:"manufacturer_id" validate:"gt=0"`
	Name           string             `json:"name" validate:"gt=0,lt=255,name"`
	VendorCode     string             `json:"vendor_code" validate:"gt=0,lte=100,vendor-code"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	UpdatedAt      pgtype.Timestamptz `json:"updated_at"`
	DeletedAt      pgtype.Timestamptz `json:"deleted_at"`
}

// Convert data to proto format
func (p *Part) ToPb() *pb.Part {
	res := pb.Part{
		Id:             p.Id,
		ManufacturerId: p.ManufacturerId,
		Name:           p.Name,
		VendorCode:     p.VendorCode,
		CreatedAt:      p.CreatedAt.Time.Unix(),
		UpdatedAt:      p.UpdatedAt.Time.Unix(),
	}

	if p.DeletedAt.Status == pgtype.Present {
		res.DeletedAt = p.DeletedAt.Time.Unix()
	}

	return &res
}

// Insert a new part in the table
func Create(ctx context.Context, conn *pgxpool.Conn, in *Part) (*Part, error) {
	var out Part

	err := conn.QueryRow(ctx, `
		INSERT INTO parts (manufacturer_id, name, vendor_code)
			VALUES($1, $2, $3)
			RETURNING id, created_at
	`, in.ManufacturerId, in.Name, in.VendorCode).
		Scan(&out.Id, &out.CreatedAt)
	if err != nil {
		return nil, errors.DatabaseError.Wrap(err, "create part")
	}

	return &out, err
}

// Select a part by id
func Read(ctx context.Context, conn *pgxpool.Conn, id uint64) (*Part, error) {
	out := Part{Id: id}

	err := conn.QueryRow(ctx, `
		SELECT manufacturer_id, name, vendor_code, created_at, updated_at, deleted_at
		FROM parts
		WHERE id = $1
	`, out.Id).
		Scan(&out.ManufacturerId, &out.Name, &out.VendorCode, &out.CreatedAt, &out.UpdatedAt, &out.DeletedAt)
	if err != nil {
		return nil, errors.DatabaseError.Wrap(err, "read part")
	}

	return &out, err
}

// Update a part by id
func Update(ctx context.Context, conn *pgxpool.Conn, in *Part) (*Part, error) {
	out := Part{Id: in.Id, ManufacturerId: in.ManufacturerId, Name: in.Name, VendorCode: in.VendorCode}

	err := conn.QueryRow(ctx, `
		UPDATE parts SET manufacturer_id = $1, name = $2, vendor_code = $3
		WHERE id = $4
		RETURNING updated_at
	`, in.ManufacturerId, in.Name, in.VendorCode, in.Id).
		Scan(&out.UpdatedAt)
	if err != nil {
		return nil, errors.DatabaseError.Wrap(err, "update part")
	}

	return &out, nil
}

// Delete a part by id
func Delete(ctx context.Context, conn *pgxpool.Conn, id uint64) error {
	ct, err := conn.Exec(ctx, `
		DELETE FROM parts
		WHERE id = $1
	`, id)
	if err != nil {
		return errors.DatabaseError.Wrap(err, "delete part")
	}

	if ct.RowsAffected() == 0 {
		return errors.NotFound.Wrap(err, "part not found")
	}

	return nil
}
