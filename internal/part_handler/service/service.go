package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"part_handler/internal/part_handler/config"
	"part_handler/internal/part_handler/models/part"
	pb "part_handler/pkg/api/v1"
	"part_handler/pkg/errors"
)

type partFunctions interface {
	CreatePart(context.Context, *part.Part) error
	ReadPart(context.Context, uint64) (*part.Part, error)
	UpdatePart(context.Context, *part.Part) (*part.Part, error)
	DeletePart(context.Context, uint64) error
}

type validator interface {
	Struct(interface{}) error
}

type partServer struct {
	db       partFunctions
	validate validator

	pb.UnimplementedPartServiceServer
}

func New(db partFunctions, conf *config.Config) *partServer {
	return &partServer{db: db, validate: conf.Validator.Validate}
}

// Create a new abstract part
func (s *partServer) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateRes, error) {
	p := part.New(req.Part)
	if err := s.validate.Struct(p); err != nil {
		errors.ValidateErrors(err)
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if err := s.db.CreatePart(ctx, p); err != nil {
		return nil, err
	}

	return &pb.CreateRes{Part: p.Convert()}, nil
}

// Read a new abstract part
func (s *partServer) Read(ctx context.Context, req *pb.ReadReq) (*pb.ReadRes, error) {
	p, err := s.db.ReadPart(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.ReadRes{Part: p.Convert()}, nil
}
