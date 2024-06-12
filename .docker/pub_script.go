package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"math/big"
	"strings"
	"time"
)

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type Order struct {
	OrderUID          string   `json:"order_uid"`
	TrackNumber       string   `json:"track_number"`
	Entry             string   `json:"entry"`
	Delivery          Delivery `json:"delivery"`
	Payment           Payment  `json:"payment"`
	Items             []Item   `json:"items"`
	Locale            string   `json:"locale"`
	InternalSignature string   `json:"internal_signature"`
	CustomerID        string   `json:"customer_id"`
	DeliveryService   string   `json:"delivery_service"`
	ShardKey          string   `json:"shardkey"`
	SmID              int      `json:"sm_id"`
	DateCreated       string   `json:"date_created"`
	OofShard          string   `json:"oof_shard"`
}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			log.Fatal(err)
		}
		b[i] = charset[num.Int64()]
	}
	return string(b)
}

func randomPhone() string {
	return "+7" + strings.Repeat("0123456789", 9)
}

func randomEmail() string {
	return randomString(5) + "@gmail.com"
}

func generateJSON() (string, error) {
	orderUID := randomString(20)
	trackNumber := randomString(12)

	order := Order{
		OrderUID:    orderUID,
		TrackNumber: trackNumber,
		Entry:       "WBIL",
		Delivery: Delivery{
			Name:    "Test Testov",
			Phone:   randomPhone(),
			Zip:     strings.Repeat("0", 7),
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   randomEmail(),
		},
		Payment: Payment{
			Transaction:  orderUID,
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       randomInt(1000, 2000),
			PaymentDt:    int(time.Now().Unix()),
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   randomInt(100, 500),
			CustomFee:    0,
		},
		Items: []Item{
			{
				ChrtID:      randomInt(1000000, 9999999),
				TrackNumber: trackNumber,
				Price:       randomInt(100, 10000),
				Rid:         randomString(15),
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  randomInt(100, 100000),
				NmID:        randomInt(1000000, 9999999),
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		ShardKey:          fmt.Sprintf("%d", randomInt(1, 10)),
		SmID:              randomInt(100, 999),
		DateCreated:       time.Now().UTC().Format(time.RFC3339),
		OofShard:          fmt.Sprintf("%d", randomInt(1, 10)),
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		return "JSON generation ERROR", err
	}

	return string(orderJSON), nil
}

func randomInt(min, max int) int {
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		log.Fatal(err)
	}
	return int(num.Int64()) + min
}

func main() {

	sc, err := stan.Connect("test-cluster", "client-pub", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	for i := 0; i < 10; i++ {
		generatedJSON, err := generateJSON()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		err = sc.Publish("Orders", []byte(generatedJSON))
		if err != nil {
			log.Fatal(err)
		}
	}
}
