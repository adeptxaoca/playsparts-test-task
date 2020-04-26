package main

import (
	"context"
	"flag"
	"os"

	"part_handler/internal/app/config"
	"part_handler/internal/pkg/logger"
	"part_handler/internal/pkg/server"
)

var (
	configPath = flag.String("config", "configs", "config file path")
	logPath    = flag.String("log", "logs/parts.log", "log file path")
)

func main() {
	flag.Parse()

	log := logger.InitLogger(*logPath)
	defer func() { _ = log.Sync() }()

	conf, err := config.AppConfiguration(*configPath)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	if err := server.Run(context.Background(), log, conf); err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
