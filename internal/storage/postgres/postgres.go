package postgres

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	wilb_l0 "wbtech-l0"
)

type Storage struct {
	Db *pgxpool.Pool
}

func (s *Storage) SaveOrder(byteValue []byte) error {

	var order wilb_l0.Order
	err := json.Unmarshal(byteValue, &order)
	if err != nil {
		return err
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}

	_, err = s.Db.Exec(context.Background(), `INSERT INTO orders (data) VALUES ($1)`, orderJSON)
	if err != nil {
		return err
	}

	log.Println("Order saved successfully with JSONB format")
	return nil
}
func (s *Storage) GetOrder(orderUID string) (string, error) {

	query := `SELECT data FROM orders WHERE data->>'order_uid' = $1`
	var jsonData string

	err := s.Db.QueryRow(context.Background(), query, orderUID).Scan(&jsonData)
	if err != nil {
		return "", err
	}
	return jsonData, nil
}
