package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"wbtech-l0/internal/config"
	"wbtech-l0/internal/storage/postgres"
	"wbtech-l0/nats"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}
	//log.Info("Starting http server")
	//log.Debug("debug messages are enabled")

	dbpool, err := config.NewPostgresDB(&config.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("Failed to init storage: %v", err)
		os.Exit(1)
	}
	storage := postgres.Storage{Db: dbpool}

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
