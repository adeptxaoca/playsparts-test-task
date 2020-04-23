package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jackc/pgx/v4"
	pb "part_handler/pkg/api/v1"
)

type partServer struct {
	db *pgx.Conn

	pb.UnimplementedPartServiceServer
}

func New(conn *pgx.Conn) *partServer {
	return &partServer{db: conn}
}

func (s *partServer) Create(ctx context.Context, req *pb.CreateReq) (*pb.CreateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
