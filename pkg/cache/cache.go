package cache

import (
	"encoding/json"
	"errors"
	"log"
	wilb_l0 "wbtech-l0"
	"wbtech-l0/pkg/storage/postgres"
)

func SaveToCache(byteValue []byte, CacheStorage map[string]wilb_l0.Order) error {
	var order wilb_l0.Order
	err := json.Unmarshal(byteValue, &order)
	if err != nil {
		return err
	}

	CacheStorage[order.OrderUID] = order
	return nil
}

func LoadOrdersToCache(storage *postgres.Storage, CacheStorage map[string]wilb_l0.Order) error {
	orders, err := storage.GetAllOrders()
	if err != nil {
		return err
	}

	for _, jsonData := range orders {
		var order wilb_l0.Order
		err := json.Unmarshal([]byte(jsonData), &order)
		if err != nil {
			log.Printf("Error unmarshaling order: %v", err)
			continue
		}

		CacheStorage[order.OrderUID] = order
	}

	return nil
}

func GetOrderFromCache(orderUID string, CacheStorage map[string]wilb_l0.Order) (wilb_l0.Order, error) {
	order, err := CacheStorage[orderUID]
	if !err {
		log.Printf("Order with OrderUID %s not found in cache", orderUID)
		return wilb_l0.Order{}, errors.New("Internal Error")
	}
	return order, nil
}
