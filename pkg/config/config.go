package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type NatsCfg struct {
	Url       string
	SubName   string
	ClusterId string
	ClientId  string
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type HttpCfg struct {
	Host string
	Port string
}

func NewPostgresDB(cfg *DBConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	dbpool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		dbpool.Close()
		return nil, err
	}

	return dbpool, nil
}
