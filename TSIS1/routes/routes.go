package routes

import (
	"net/http"
	"webapp/handlers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health-check", healthCheckHandler).Methods("GET")
	r.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{title}", handlers.GetBookByTitle).Methods("GET")
	return r
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Welcome to my Book API! - By Your Name"))
}
