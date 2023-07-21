package event

type EventRepositoryInterface interface {
	Save(event *Event) error
	FindAll() ([]*Event, error)
}
