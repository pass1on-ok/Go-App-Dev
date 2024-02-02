package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tsis1-aot-api/api"

	"github.com/gorilla/mux"
)

func generateHTML(t api.Titan) string {
	abilitiesList := ""
	for _, ability := range t.Abilities {
		abilitiesList += fmt.Sprintf("<li>%s</li>\n", ability)
	}
	html := fmt.Sprintf(`<img src="%s">
	<h1>%s</h1>
	<p><strong>Height:</strong>%d</p>
	<h2>Abilities:</h2>
	<ul>
	%s
	</ul>
	<p><strong>Ally:</strong>%s</p>`,
		t.Image, t.Name, t.Height, abilitiesList, t.Ally)

	return html
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to AOT Api made by Uldana🛡️")
}

func Titans(w http.ResponseWriter, r *http.Request) {
	log.Println("entering titans end point")
	titans := api.Titans

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(titans)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
func Titan(w http.ResponseWriter, r *http.Request) {
	log.Println("entering titan end point")
	w.Header().Set("Content-Type", "text/html")
	titans := api.Titans
	vars := mux.Vars(r)
	i, err := strconv.Atoi(vars["titan"])
	if err != nil || i < 1 || i > 9 {
		fmt.Fprintf(w, `<h1>404 page not found</h1>`)
		return
	}
	titan := titans[i-1]

	fmt.Fprintf(w, generateHTML(titan))

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// jsonResponse, err := json.Marshal(titans)
	// if err != nil {
	// 	return
	// }
	// w.Write(jsonResponse)
}
