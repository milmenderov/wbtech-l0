package main

import (
	stan "github.com/nats-io/stan.go"
	"log"
)

func main() {
	// Подключение к серверу с указанием адреса и порта
	sc, err := stan.Connect("test-cluster", "client-pub", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	// Публикация сообщения
	err = sc.Publish("foo", []byte("Hello, NATS Streaming with address and port!"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Сообщение отправлено!")
}
