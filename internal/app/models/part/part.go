package part

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"

	pb "part_handler/internal/pkg/api/v1"
	"part_handler/internal/pkg/errors"
	"part_handler/internal/pkg/utils"
)

type Part struct {
	Id             uint64             `json:"id"`
	ManufacturerId uint64             `json:"manufacturer_id"`
	Name           string             `json:"name" validate:"omitempty,gt=0,lt=255,name"`
	VendorCode     string             `json:"vendor_code" validate:"omitempty,gt=0,lte=100,vendor-code"`
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
	}

	if p.CreatedAt.Status == pgtype.Present {
		res.CreatedAt = p.CreatedAt.Time.Unix()
	}

	if p.UpdatedAt.Status == pgtype.Present {
		res.UpdatedAt = p.UpdatedAt.Time.Unix()
	}

	if p.DeletedAt.Status == pgtype.Present {
		res.DeletedAt = p.DeletedAt.Time.Unix()
	}

	return &res
}

// Insert a new part in the table
func Create(ctx context.Context, conn *pgxpool.Conn, in *Part) (*Part, error) {
	out := Part{ManufacturerId: in.ManufacturerId, Name: in.Name, VendorCode: in.VendorCode}

	err := conn.QueryRow(ctx, `
		INSERT INTO parts (manufacturer_id, name, vendor_code)
			VALUES($1, $2, $3)
			RETURNING id, created_at;
	`, in.ManufacturerId, in.Name, in.VendorCode).
		Scan(&out.Id, &out.CreatedAt)

	return &out, err
}

// Select a part by id
func Read(ctx context.Context, conn *pgxpool.Conn, id uint64) (*Part, error) {
	out := Part{Id: id}

	err := conn.QueryRow(ctx, `
		SELECT manufacturer_id, name, vendor_code, created_at, updated_at, deleted_at
		FROM parts
		WHERE id = $1;
	`, id).
		Scan(&out.ManufacturerId, &out.Name, &out.VendorCode, &out.CreatedAt, &out.UpdatedAt, &out.DeletedAt)

	return &out, err
}

// Update a part by id
func Update(ctx context.Context, conn *pgxpool.Conn, in *Part) (*Part, error) {
	out := Part{Id: in.Id}

	set := updatePrepare(in)
	if len(set) == 0 {
		return nil, errors.ValidationError.New("need one or more 'part' fields")
	}

	err := conn.QueryRow(ctx, `
		UPDATE parts SET `+set+`
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING manufacturer_id, name, vendor_code, created_at, updated_at;
	`, in.Id).
		Scan(&out.ManufacturerId, &out.Name, &out.VendorCode, &out.CreatedAt, &out.UpdatedAt)

	return &out, err
}

// Delete a part by id
func Delete(ctx context.Context, conn *pgxpool.Conn, id uint64) error {
	_, err := conn.Exec(ctx, `
		UPDATE parts SET deleted_at = $2
		WHERE id = $1 AND deleted_at IS NULL;
	`, id, time.Now())

	return err
}

func updatePrepare(in *Part) string {
	var set []string

	if in.ManufacturerId > 0 {
		set = append(set, fmt.Sprintf("manufacturer_id = '%d'", in.ManufacturerId))
	}
	if in.Name != "" {
		set = append(set, fmt.Sprintf("name = %s", utils.QuoteString(in.Name)))
	}
	if in.VendorCode != "" {
		set = append(set, fmt.Sprintf("vendor_code = %s", utils.QuoteString(in.VendorCode)))
	}
	if len(set) > 0 {
		set = append(set, fmt.Sprintf("updated_at = '%s'",
			time.Now().Truncate(time.Microsecond).Format("2006-01-02 15:04:05.999999999Z07:00:00")))
	}

	return strings.Join(set, ",")
}
