package event

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"io"
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

func PushEvent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var e Event
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// Handle the error (e.g., return an error response)
		return
	}

	err = json.Unmarshal(body, &e)
	if err != nil {
		// Handle the error (e.g., return an error response)
		return
	}
	e.ID = genId()
	e.Time = time.Now()

	//event, _ := NewEvent(genId(), "name", "origin", "description")
	err = e.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		//_ = json.NewEncoder(w).Encode(err.Error())
	} else {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(e)

	}
}

func genId() string {
	return uuid.New().String()
}
