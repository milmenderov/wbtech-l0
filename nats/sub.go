package main

import (
	stan "github.com/nats-io/stan.go"
	"log"
)

func main() {
	// Подключение к серверу с указанием адреса и порта
	sc, err := stan.Connect("test-cluster", "client-sub", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	// Подписка на канал
	_, err = sc.Subscribe("foo", func(m *stan.Msg) {
		log.Printf("Получено сообщение: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// Ожидание завершения
	select {}
}
