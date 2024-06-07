package nats

import (
	stan "github.com/nats-io/stan.go"
	"log"
	"wbtech-l0/internal/storage/postgres"
)

func NatsConnect(storage postgres.Storage) error {

	sc, err := stan.Connect("test-cluster", "client-sub", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	_, err = sc.Subscribe("Orders", func(m *stan.Msg) {
		log.Printf("Получено сообщение: %s\n", string(m.Data))

		err := storage.SaveOrder(m.Data)
		if err != nil {
			log.Printf("Ошибка при сохранении: %v", err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	select {}
}
