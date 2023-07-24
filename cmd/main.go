package main

import (
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
	"github/fpinna/event-time-line/pkg/rabbitmq"
	"github/fpinna/event-time-line/pkg/router"
	"log"
	"net/http"
)

const (
	queueName = "events"
)

func main() {
	var err error

	// RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	messagingApp, err := rabbitmq.NewRabbitMQService(conn, queueName)
	messagingService := rabbitmq.NewMessagingService(messagingApp)

	r := router.SetupRouter(messagingService)

	// Consumer Test - RabbitMQ
	receivedMessages, err := messagingService.ReceiveMessages(queueName)
	go func() {
		for d := range receivedMessages {
			log.Printf("Received a message: %s", d.Body)
		}

	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	// End Consumer

	// Server start
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
