package event

import (
	message "github/fpinna/event-time-line/pkg/rabbitmq"
)

type EventInterfaceMessage interface {
	message.MessagingService
}
