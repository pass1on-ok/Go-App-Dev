// cmd/handlers/main.go
package main

import (
	"log"
	"net/http"
	"webapp/pkg/routes"
)

func main() {
	r := routes.NewRouter()

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
