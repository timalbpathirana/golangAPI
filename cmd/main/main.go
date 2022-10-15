package main

import (
	"github.com/gorilla/mux"
	"github.com/timalbpathirana/golangAPI/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// Registering the routes
	routes.RegisterStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))

}
