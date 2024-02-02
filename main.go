// main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/my-web-app/routes"
)

func main() {
	r := routes.NewRouter()

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
