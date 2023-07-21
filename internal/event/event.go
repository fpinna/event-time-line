package event

import (
	"encoding/json"
	//"github.com/fpinna/event-time-line/internal/entity"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"time"
)

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

func (ev *Event) PushEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	event, _ := NewEvent(genId(), "name", "origin", "description")
	err := event.Validate()
	if err != nil {
		_ = json.NewEncoder(w).Encode(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(event)

}

func genId() string {
	return uuid.New().String()
}
