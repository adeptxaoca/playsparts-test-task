package main

import (
	"context"
	"flag"
	"os"

	"go.uber.org/zap"

	"part_handler/internal/app/config"
	"part_handler/internal/app/server"
)

var (
	port = flag.Uint("port", 3000, "grpc server port")
)

func main() {
	flag.Parse()

	log, _ := zap.NewProduction()
	defer func() { _ = log.Sync() }()

	conf, err := config.AppConfiguration()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	if err := server.Run(context.Background(), *port, log, conf); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
