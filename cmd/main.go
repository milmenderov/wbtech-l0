package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"wbtech-l0/cache"
	"wbtech-l0/internal/config"
	"wbtech-l0/internal/storage/postgres"
	"wbtech-l0/nats"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	dbpool, err := config.NewPostgresDB(&config.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	storage := &postgres.Storage{Db: dbpool}

	if err != nil {
		log.Fatalf("Failed to init storage: %v", err)
		os.Exit(1)
	}

	err = cache.LoadOrdersToCache(storage)
	if err != nil {
		log.Fatalf("Error loading orders to cache: %v", err)
	}

	//cache.PrintCache()

	err = nats.NatsConnect(storage)
	if err != nil {
		log.Fatalf("Error connect to NATS: %v", err)
	}

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
