package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jessicasomaiya/sudoku/cmd/api/server/handlers"
)

var staticDir = "cmd/api/server/static"

// chi handler?

var (
	LOOPS = 10000
	SIZE  = 9
)

func main() {

	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", handlers.Hello)
	http.HandleFunc("/generate", handlers.Run)

	fmt.Print("\n✨Starting server at port 8080.✨\n\nClick here 👉: http://localhost:8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
