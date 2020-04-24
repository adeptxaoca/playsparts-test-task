package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"part_handler/internal/part_handler/config"
	"part_handler/internal/part_handler/database"
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

	db, err := database.Connect(context.Background(), conf)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer func() { _ = db.Conn.Close(context.Background()) }()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", conf.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPartServiceServer(grpcServer, server.New(db, conf))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
