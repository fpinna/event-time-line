// application
package rabbitmq

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessagingService interface {
	SendMessage(msg Message, queueName string) error
	ReceiveMessages(queueName string) (<-chan amqp.Delivery, error)
}

type rabbitMQService struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQService(conn *amqp.Connection, queueName string) (MessagingService, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &rabbitMQService{
		conn: conn,
		ch:   ch,
	}, nil
}

func (s *rabbitMQService) SendMessage(msg Message, queueName string) error {
	body := []byte(msg.Body)
	err := s.ch.PublishWithContext(context.Background(), "", queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *rabbitMQService) ReceiveMessages(queueName string) (<-chan amqp.Delivery, error) {
	_, err := s.ch.QueueDeclare(queueName, false, false, false, false, nil)
	msgs, err := s.ch.Consume(queueName, fmt.Sprintf("%s-consumer", queueName), true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (s *rabbitMQService) Close() {
	s.ch.Close()
}
