// pkg/routes/routes.go
package routes

import (
	"encoding/json"
	"net/http"
	"webapp/pkg/models"

	"github.com/gorilla/mux"
)

var books = []models.Book{
	{Title: "The Hobbit", Author: "J.R.R. Tolkien", Year: 1937},
	{Title: "Harry Potter and the Sorcerer's Stone", Author: "J.K. Rowling", Year: 1997},
	{Title: "To Kill a Mockingbird", Author: "Harper Lee", Year: 1960},
	{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925},
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health-check", HealthCheckHandler).Methods("GET")
	r.HandleFunc("/books", GetAllBooks).Methods("GET")
	r.HandleFunc("/books/{title}", GetBookByTitle).Methods("GET")
	return r
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Welcome to my Book API! \nCreated by Kumar"))
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBookByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	title := params["title"]

	for _, book := range books {
		if book.Title == title {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}
