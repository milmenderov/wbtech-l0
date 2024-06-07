package main

import (
	"log"
	"os"
	"wbtech-l0/internal/config"
	"wbtech-l0/internal/storage/postgres"
	"wbtech-l0/nats"
)

func main() {
	cfg := config.LoadConfig()

	dbpool, err := config.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to init storage: %v", err)
		os.Exit(1)
	}
	storage := postgres.Storage{Db: dbpool}

	err = nats.NatsConnect(storage)
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS: %v", err)
	}
}
