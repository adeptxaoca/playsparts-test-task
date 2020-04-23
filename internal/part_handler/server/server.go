package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "part_handler/pkg/api/v1"
)

type partServer struct {
	pb.UnimplementedPartServiceServer
}

func New() *partServer {
	return &partServer{}
}

func (s *partServer) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
