package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"part_handler/internal/part_handler/server"
	pb "part_handler/pkg/api/v1"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPartServiceServer(grpcServer, server.New())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
