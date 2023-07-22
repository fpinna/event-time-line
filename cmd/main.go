package main

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
	"github/fpinna/event-time-line/internal/event"
	"github/fpinna/event-time-line/pkg/rabbitmq"
	"net/http"
)

func main() {
	// RabbitMQ
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"events",
		false,
		false,
		false,
		false,
		nil,
	)

	//msgRabbitmqChannel := make(chan amqp.Delivery)
	//err = rabbitmq.Publish(ch, amqp.Delivery{Body: []byte("hola")})
	//if err != nil {
	//	panic(err)
	//}

	// New database
	//dbf, err := sql.Open("sqlite3", "./events.db")
	//if err != nil {
	//	panic(err)
	//}
	//defer dbf.Close()
	// New repository
	//db.NewEventRepository(dbf)
	// New router
	router := chi.NewRouter()

	// Middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("status"))
		if err != nil {
			return
		}
	})

	router.Route("/event", func(r chi.Router) {

		r.Post("/push", event.PushEvent)
		err = rabbitmq.PublishEvent(ch, amqp.Delivery{Body: []byte("hola")})
		if err != nil {
			panic(err)
		}
		//r.Post("/pull", func(w http.ResponseWriter, r *http.Request) {
		//	w.Write([]byte("pull"))
		//})
	})

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}

func rabbitmqProducer(ch *amqp.Channel, msg amqp.Delivery) {
	err := ch.Publish(
		"",
		"events",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg.Body,
		})
	if err != nil {
		panic(err)
	}

}
