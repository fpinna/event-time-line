package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

const QueueName = "events"

type RabbitMQ struct {
	Ch *amqp.Channel
}

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	var ch RabbitMQ
	if err != nil {
		return nil, err
	}
	ch.Ch, err = conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch.Ch, nil
}

//func Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
//	msgs, err := ch.Consume(
//		"events",
//		"go-consumer",
//		false,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		return err
//	}
//	for msg := range msgs {
//		out <- msg
//	}
//	return nil
//}

func PublishEvent(ch *amqp.Channel, msg amqp.Delivery) error {
	err := ch.Publish(
		"",
		QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg.Body,
		})
	if err != nil {
		return err
	}
	return nil
}
