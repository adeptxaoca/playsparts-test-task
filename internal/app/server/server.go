package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"part_handler/internal/app/config"
	"part_handler/internal/app/database"
	pb "part_handler/internal/pkg/api/v1"
	"part_handler/internal/pkg/errors"
	v1 "part_handler/internal/pkg/service/v1"
)

// Run gRPC service to publish Parts service
func Run(ctx context.Context, port uint, log *zap.Logger, conf *config.Config) error {
	// Connection setup and database connection
	db, ver, err := database.Setup(ctx, &conf.Database)
	if err != nil {
		return errors.Wrap(err, "Unable to setup database")
	}
	defer db.Pool.Close()

	log.Info("Migration done. Current schema version", zap.Int32("version", ver))

	// Listen announces on the local network address.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrap(err, "Failed to listen")
	}

	// Register service
	grpc_zap.ReplaceGrpcLoggerV2(log)
	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor),
			),
			grpc_zap.UnaryServerInterceptor(log),
		),
	)
	pb.RegisterPartServiceServer(grpcServer, v1.NewService(db, conf))

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Info("shutting down gRPC server...")
		grpcServer.GracefulStop()
		<-ctx.Done()
	}()

	// Start gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		return errors.Wrap(err, "Failed to serve")
	}

	return nil
}
