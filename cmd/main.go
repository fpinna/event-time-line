package main

import (
	"database/sql"
	"github/fpinna/event-time-line/internal/event"
	db "github/fpinna/event-time-line/internal/infra/database"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {

	// New database
	dbf, err := sql.Open("sqlite3", "./events.db")
	if err != nil {
		panic(err)
	}
	defer dbf.Close().Error()
	// New repository
	db.NewEventRepository(dbf)
	// New router
	router := chi.NewRouter()

	// Middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("status"))
	})

	router.Route("/event", func(r chi.Router) {
		r.Post("/push", event.PushEvent)
		r.Post("/pull", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pull"))
		})
	})

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

}
