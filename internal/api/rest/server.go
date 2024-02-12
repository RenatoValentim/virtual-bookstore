package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func StartServer(router *chi.Mux) {
	port := viper.GetInt("server_port")
	log.Printf("API listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
