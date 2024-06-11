package nats

import (
	stan "github.com/nats-io/stan.go"
	"log"
	"wbtech-l0/cache"
	"wbtech-l0/internal/storage/postgres"
)

func NatsConnect(storage *postgres.Storage) error {

	sc, err := stan.Connect("test-cluster", "client-sub", stan.NatsURL("nats://nats-streaming:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	_, err = sc.Subscribe("Orders", func(m *stan.Msg) {
		log.Printf("Message has been received: %s\n", string(m.Data))

		err := cache.SaveToCache(m.Data)
		if err != nil {
			log.Printf("Error saving to cache: %v", err)
			return
		}
		err = storage.SaveOrder(m.Data)
		if err != nil {
			log.Printf("Error saving to db: %v", err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
