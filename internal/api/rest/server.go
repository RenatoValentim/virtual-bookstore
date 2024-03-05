package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RenatoValentim/virtual-bookstore/internal/constants/environments"
	"github.com/spf13/viper"
)

func StartServer(router http.Handler) {
	port := viper.GetInt(environments.ServerPort)
	log.Printf("API listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
