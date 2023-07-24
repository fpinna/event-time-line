// interfaces/messaging.go
package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type MessagingInterface interface {
	SendMessage(body string, queueName string) error
	ReceiveMessages(queueName string) (<-chan amqp.Delivery, error)
}

type MessagingServiceStruct struct {
	messagingApp MessagingService
}

func NewMessagingService(messagingApp MessagingService) *MessagingServiceStruct {
	return &MessagingServiceStruct{
		messagingApp: messagingApp,
	}
}

func (s *MessagingServiceStruct) SendMessage(body string, queueName string) error {
	msg := Message{Body: body}
	return s.messagingApp.SendMessage(msg, queueName)
}

func (s *MessagingServiceStruct) ReceiveMessages(queueName string) (<-chan amqp.Delivery, error) {
	messages, err := s.messagingApp.ReceiveMessages(queueName)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
