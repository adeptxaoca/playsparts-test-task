package v1

import (
	"context"

	"part_handler/internal/app/config"
	"part_handler/internal/app/models/part"
	pb "part_handler/internal/pkg/api/v1"
	"part_handler/internal/pkg/errors"
)

type partFunctions interface {
	CreatePart(context.Context, *part.Part) (*part.Part, error)
	ReadPart(context.Context, uint64) (*part.Part, error)
	UpdatePart(context.Context, *part.Part) (*part.Part, error)
	DeletePart(context.Context, uint64) error
}

type validator interface {
	Struct(interface{}) error
}

type service struct {
	db       partFunctions
	validate validator

	pb.UnimplementedPartServiceServer
}

// NewService creates Parts service
func NewService(db partFunctions, conf *config.Config) *service {
	return &service{db: db, validate: conf.Validator}
}

// Create a new abstract part
func (s *service) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateRes, error) {
	in := part.Part{
		ManufacturerId: req.Part.ManufacturerId,
		Name:           req.Part.Name,
		VendorCode:     req.Part.VendorCode,
	}
	if err := s.validate.Struct(&in); err != nil {
		return nil, errors.GrpcError(err)
	}

	out, err := s.db.CreatePart(ctx, &in)
	if err != nil {
		return nil, errors.GrpcError(err)
	}

	return &pb.CreateRes{Part: out.ToPb()}, nil
}

// Read a abstract part
func (s *service) Read(ctx context.Context, req *pb.ReadReq) (*pb.ReadRes, error) {
	if req.Id == 0 {
		return nil, errors.GrpcError(errors.ValidationError.New("id must be more 0"))
	}

	out, err := s.db.ReadPart(ctx, req.Id)
	if err != nil {
		return nil, errors.GrpcError(err)
	}

	return &pb.ReadRes{Part: out.ToPb()}, nil
}

// Update a abstract part
func (s *service) Update(ctx context.Context, req *pb.UpdateReq) (*pb.UpdateRes, error) {
	in := part.Part{
		Id:             req.Part.Id,
		ManufacturerId: req.Part.ManufacturerId,
		Name:           req.Part.Name,
		VendorCode:     req.Part.VendorCode,
	}
	if err := s.validate.Struct(&in); err != nil {
		return nil, errors.GrpcError(err)
	}

	out, err := s.db.UpdatePart(ctx, &in)
	if err != nil {
		return nil, errors.GrpcError(err)
	}

	return &pb.UpdateRes{Part: out.ToPb()}, nil
}

// Delete a abstract part
func (s *service) Delete(ctx context.Context, req *pb.DeleteReq) (*pb.DeleteRes, error) {
	if req.Id == 0 {
		return nil, errors.GrpcError(errors.ValidationError.New("id must be more 0"))
	}

	err := s.db.DeletePart(ctx, req.Id)
	if err != nil {
		return nil, errors.GrpcError(err)
	}

	return &pb.DeleteRes{Success: true}, nil
}
