package nats

import (
	stan "github.com/nats-io/stan.go"
	"log"
	wilb_l0 "wbtech-l0"
	"wbtech-l0/pkg/cache"
	"wbtech-l0/pkg/config"
	"wbtech-l0/pkg/storage/postgres"
)

func NatsConnectStart(storage *postgres.Storage, NatsCfg config.NatsCfg, Cache map[string]wilb_l0.Order) error {

	sc, err := stan.Connect(NatsCfg.ClusterId, NatsCfg.ClientId, stan.NatsURL(NatsCfg.Url))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()
	log.Println("NATS connected")

	_, err = sc.Subscribe(NatsCfg.SubName, func(m *stan.Msg) {
		log.Printf("Message has been received: %s\n", string(m.Data))

		err := cache.SaveToCache(m.Data, Cache)
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
