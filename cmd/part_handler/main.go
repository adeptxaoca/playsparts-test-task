package main

import (
	"context"
	"flag"
	"log"

	"part_handler/internal/app/config"
	"part_handler/internal/pkg/server"
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

	if err := server.Run(context.Background(), conf); err != nil {
		log.Fatal(err)
	}
}
