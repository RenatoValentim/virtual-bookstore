package main

import (
	"github.com/RenatoValentim/virtual-bookstore/internal/api/rest"
	"github.com/RenatoValentim/virtual-bookstore/internal/config"
)

func main() {
	config.LoadConfig()
	routes := rest.LoadRoutes()
	rest.StartServer(routes)
}
