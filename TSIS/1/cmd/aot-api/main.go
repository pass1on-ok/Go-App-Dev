package main

import (
	"log"
	"net/http"

	"tsis1-aot-api/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting API server")
	router := mux.NewRouter()
	log.Println("creating routes")

	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/titans", handlers.Titans).Methods("GET")
	router.HandleFunc("/titans/{titan}", handlers.Titan).Methods("GET")
	http.Handle("/", router)

	http.ListenAndServe(":8080", router)

}
