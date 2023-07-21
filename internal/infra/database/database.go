package database

import (
	"database/sql"
	"github/fpinna/event-time-line/internal/event"
)

type EventRepository struct {
	Db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{Db: db}
}

func (ev *EventRepository) Save(event *event.Event) error {
	_, err := ev.Db.Exec("INSERT INTO events (time ,id, name, origin, description) VALUES ($1, $2, $3, $4, $5)", event.Time, event.ID, event.Name, event.Origin, event.Description)
	if err != nil {
		return err
	}
	return nil
}

func (ev *EventRepository) FindAll() ([]*event.Event, error) {
	rows, err := ev.Db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*event.Event
	for rows.Next() {
		var e event.Event
		err := rows.Scan(&e.Time, &e.ID, &e.Name, &e.Origin, &e.Description)
		if err != nil {
			return nil, err
		}
		events = append(events, &e)
	}
	return events, nil
}
