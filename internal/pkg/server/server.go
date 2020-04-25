package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"part_handler/internal/app/config"
	"part_handler/internal/app/database"
	pb "part_handler/internal/pkg/api/v1"
	"part_handler/internal/pkg/errors"
	v1 "part_handler/internal/pkg/service/v1"
)

// Run gRPC service to publish Parts service
func Run(ctx context.Context, conf *config.Config) error {
	// Connection setup and database connection
	db, err := database.Setup(ctx, conf)
	if err != nil {
		return errors.Wrap(err, "Unable to setup database")
	}
	defer db.Pool.Close()

	// Listen announces on the local network address.
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", conf.Server.Port))
	if err != nil {
		return errors.Wrap(err, "Failed to listen")
	}

	// Register service
	grpcServer := grpc.NewServer()
	pb.RegisterPartServiceServer(grpcServer, v1.NewService(db, conf))

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
		return errors.Wrap(err, "Failed to serve")
	}

	return nil
}
