package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

func LoadRoutes() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ping", ping)

	r.Route("/author", func(r chi.Router) {
		r.Post("/register", registerAuthorHandler)
	})

	port := 8000
	log.Printf("API listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), r))
}
