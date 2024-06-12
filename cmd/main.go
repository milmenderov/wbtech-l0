package main

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	wilb_l0 "wbtech-l0"
	"wbtech-l0/pkg/cache"
	"wbtech-l0/pkg/config"
	"wbtech-l0/pkg/http/handlers"
	"wbtech-l0/pkg/http/service"
	"wbtech-l0/pkg/nats"
	"wbtech-l0/pkg/storage/postgres"
)

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func StartHttpServer(HttpCfg config.HttpCfg, CacheStorage map[string]wilb_l0.Order) error {
	services := service.NewService()
	handler := handlers.NewHandler(services, CacheStorage)
	mux := handler.InitRoutes()
	return http.ListenAndServe(HttpCfg.Host+":"+HttpCfg.Port, mux)
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}
	log.Println("Config initialized")

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
	log.Println("Storage initialized")

	var CacheStorage = make(map[string]wilb_l0.Order)
	err = cache.LoadOrdersToCache(storage, CacheStorage)
	if err != nil {
		log.Fatalf("Error loading orders to cache: %v", err)
	}
	log.Println("Cache initialized")

	HttpCfg := config.HttpCfg{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT"),
	}
	go StartHttpServer(HttpCfg, CacheStorage)
	log.Println("HTTP server started")

	NatsCfg := config.NatsCfg{
		Url:       os.Getenv("NATS_URL"),
		SubName:   os.Getenv("NATS_SUB_NAME"),
		ClusterId: os.Getenv("NATS_CLUSTER_ID"),
		ClientId:  os.Getenv("NATS_CLIENT_ID"),
	}
	err = nats.NatsConnectStart(storage, NatsCfg, CacheStorage)

	if err != nil {
		log.Fatalf("Error connect to NATS: %v", err)
	}

}
