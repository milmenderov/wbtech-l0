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

func (s *Storage) GetAllOrders() ([]string, error) {
	query := `SELECT data FROM orders`
	rows, err := s.Db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []string
	for rows.Next() {
		var jsonData string
		if err := rows.Scan(&jsonData); err != nil {
			return nil, err
		}
		orders = append(orders, jsonData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
