package event

type EventRepositoryInterface interface {
	Save(event *Event) error
	FindAll() ([]*Event, error)
	PublishEvent(event *Event) error
	PushEvent(event *Event) error
}
