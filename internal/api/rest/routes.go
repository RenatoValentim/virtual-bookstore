package rest

import (
	"encoding/json"
	"net/http"

	"github.com/RenatoValentim/virtual-bookstore/internal/api/rest/handlers"
	"github.com/RenatoValentim/virtual-bookstore/internal/api/rest/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func LoadRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", ping)

	r.Route("/author", func(r chi.Router) {
		r.Use(middlewares.AuthorInput)

		r.Post("/register", handlers.RegisterAuthorHandler)
	})

	return r
}
