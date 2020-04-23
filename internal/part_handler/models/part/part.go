package part

import (
	"time"

	pb "part_handler/pkg/api/v1"
)

type Part struct {
	Id             uint64    `json:"id"`
	ManufacturerId uint64    `json:"manufacturer_id" validate:"gt=0"`
	Name           string    `json:"name" validate:"gt=0,lt=255,name"`
	VendorCode     string    `json:"vendor_code" validate:"gt=0,lte=100,vendor-code"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

func New(p *pb.Part) *Part {
	return &Part{
		ManufacturerId: p.ManufacturerId,
		Name:           p.Name,
		VendorCode:     p.VendorCode,
	}
}

func (p *Part) Convert() *pb.Part {
	return &pb.Part{
		Id:             p.Id,
		ManufacturerId: p.ManufacturerId,
		Name:           p.Name,
		VendorCode:     p.VendorCode,
		CreatedAt:      p.CreatedAt.Unix(),
		UpdatedAt:      p.UpdatedAt.Unix(),
		DeletedAt:      p.DeletedAt.Unix(),
	}
}
