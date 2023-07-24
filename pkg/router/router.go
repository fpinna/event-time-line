package router

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	eventmsg "github/fpinna/event-time-line/internal/event"
	interfaces "github/fpinna/event-time-line/pkg/rabbitmq"
	"net/http"
	"time"
)

func SetupRouter(messagingService interfaces.MessagingInterface) http.Handler {
	r := chi.NewRouter()
	// Middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/push", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var event eventmsg.Event
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			http.Error(w, "Erro ao decodificar o JSON", http.StatusBadRequest)
			return
		}

		event.ID = uuid.New().String()
		event.Time = time.Now()

		err = event.Validate()

		err = messagingService.SendMessage(fmt.Sprintf("%s", event), "events")
		if err != nil {
			http.Error(w, "Erro ao enviar a mensagem", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		//_ = json.NewEncoder(w).Encode(event)
	})

	return r
}
