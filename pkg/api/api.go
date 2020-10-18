package api

import (
	"fmt"
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/", sudokuHome)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

func sudokuHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}
