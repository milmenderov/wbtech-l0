package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"log"
	"os"
	"wbtech-l0/internal/storage/postgres"
)

func main() {

	jsonFile, err := os.Open("contract.json")
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	dsn := "postgres://postgres:qwerty@localhost:5432/postgres"
	dbpool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	storage := postgres.Storage{Db: dbpool}

	if err := storage.SaveOrder(byteValue); err != nil {
		log.Fatalf("Error saving order to database: %v", err)
	}

}
