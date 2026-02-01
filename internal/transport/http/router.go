package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"thought-RestAPI/internal/transport/http/handler"
	"thought-RestAPI/internal/usecase"
)

func Router(service *usecase.Thought) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Get("/ready", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })

	handlers := handler.NewHandlers(service)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/thoughts", func(r chi.Router) {
				r.Get("/random", handlers.GetRandomThought)
				r.Post("/", handlers.CreateThought)
				r.Put("/{id}", handlers.UpdateThought)
				r.Delete("/{id}", handlers.DeleteThought)
			})
		})
	})

	return r
}
