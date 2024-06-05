package postgres

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	db *sql.DB
}

func (s *Storage) SaveMessage() error {
	const op = "storage.postgres.SaveURL"
	// Парсинг JSON данных
	var order Order
	err = json.Unmarshal(byteValue, &order)
	if err != nil {
		log.Fatal(err)
	}

	// Вставка данных в таблицу orderr
	_, err = db.Exec(`INSERT INTO orderr (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) 
					  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated, order.OofShard)
	if err != nil {
		log.Fatal(err)
	}

	// Вставка данных в таблицу delivery
	_, err = db.Exec(`INSERT INTO delivery (name, phone, zip, city, address, region, email) 
					  VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		log.Fatal(err)
	}

	// Вставка данных в таблицу payment
	_, err = db.Exec(`INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) 
					  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		log.Fatal(err)
	}

	// Вставка данных в таблицу items
	for _, item := range order.Items {
		_, err = db.Exec(`INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) 
						  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
			item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
		if err != nil {
			log.Fatal(err)
		}
	}

}
