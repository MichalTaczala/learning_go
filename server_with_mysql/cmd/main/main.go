package main

import (
	"github.com/gorilla/mux"
	"github.com/radakanis/sqlServerGo/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
}
