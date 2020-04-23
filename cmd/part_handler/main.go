package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"

	"part_handler/internal/part_handler/config"
	"part_handler/internal/part_handler/server"
	pb "part_handler/pkg/api/v1"
)

var (
	configPath = flag.String("config", "configs", "config file path")
)

func main() {
	flag.Parse()

	conf, err := config.AppConfiguration(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.Connect(context.Background(), conf.ConnString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	lis, err := net.Listen("tcp", conf.NetAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPartServiceServer(grpcServer, server.New(conn))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
