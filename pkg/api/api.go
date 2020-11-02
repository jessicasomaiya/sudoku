package api

import (
	"fmt"
	"log"
	"net/http"
)

func HandleRequests(generate func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc("/", sudokuHome)
	http.HandleFunc("/generate", generate)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

func sudokuHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func generateOne() {
}
