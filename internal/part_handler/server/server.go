package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"part_handler/internal/part_handler/config"
	"part_handler/internal/part_handler/models/part"
	pb "part_handler/pkg/api/v1"
)

type partFunctions interface {
	CreatePart(context.Context, *part.Part) (*part.Part, error)
	ReadPart(context.Context, uint64) (*part.Part, error)
}

type validator interface {
	Struct(s interface{}) error
}

type partServer struct {
	db partFunctions
	v  validator

	pb.UnimplementedPartServiceServer
}

func New(db partFunctions, conf *config.Config) *partServer {
	return &partServer{db: db, v: conf.Validator.Validate}
}

// Create a new abstract part
func (s *partServer) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateRes, error) {
	partIn := part.New(req.Part)
	if err := s.v.Struct(partIn); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	partOut, err := s.db.CreatePart(ctx, partIn)
	if err != nil {
		return nil, err
	}

	return &pb.CreateRes{Part: partOut.Convert()}, nil
}
