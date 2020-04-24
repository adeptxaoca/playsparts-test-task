package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"part_handler/internal/part_handler/config"
	"part_handler/internal/part_handler/database"
	"part_handler/internal/part_handler/service"
	pb "part_handler/pkg/api/v1"
)

var (
	configPath = flag.String("config", "configs", "config file path")
)

func main() {
	ctx := context.Background()
	flag.Parse()

	conf, err := config.AppConfiguration(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Setup(ctx, conf)
	if err != nil {
		log.Fatalf("Unable to setup database: %v\n", err)
	}
	defer db.Pool.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", conf.Server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register service
	grpcServer := grpc.NewServer()
	pb.RegisterPartServiceServer(grpcServer, service.New(db, conf))

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("shutting down gRPC server...")
			grpcServer.GracefulStop()
			<-ctx.Done()
		}
	}()

	// Start gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
