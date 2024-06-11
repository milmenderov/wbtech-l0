package cache

import (
	"encoding/json"
	"log"
	wilb_l0 "wbtech-l0"
	"wbtech-l0/internal/storage/postgres"
)

var Cache = make(map[string]wilb_l0.Order)

//func PrintCache() {
//	for key, order := range cache {
//		fmt.Printf("OrderUID: %s, Order: %+v\n", key, order)
//	}
//}

func SaveToCache(byteValue []byte) error {
	var order wilb_l0.Order
	err := json.Unmarshal(byteValue, &order)
	if err != nil {
		return err
	}

	Cache[order.OrderUID] = order
	return nil
}

func LoadOrdersToCache(storage *postgres.Storage) error {
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

		Cache[order.OrderUID] = order
	}

	return nil
}
