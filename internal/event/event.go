package event

import (
	"errors"
	"time"
)

// Message - Domain model
type Event struct {
	ID          string
	Name        string
	Origin      string
	Description string
	Time        time.Time
}

func NewEvent(id string, name string, origin string, description string) (*Event, error) {
	ev := Event{
		ID:          id,
		Name:        name,
		Origin:      origin,
		Description: description,
	}
	return &ev, nil
}

func (ev *Event) Validate() error {
	if ev.ID == "" {
		return errors.New("ID error")
	}
	if ev.Origin == "" {
		// check if in database
		return errors.New("origin is required")
	}
	if ev.Name == "" {
		return errors.New("name is required")
	}
	if ev.Description == "" {
		return errors.New("description is required")
	}
	return nil
}
