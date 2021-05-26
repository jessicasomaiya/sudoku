package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jessicasomaiya/sudoku/pkg/solver"
)

var staticDir = "./pkg/server/static"

// chi handler?

func New() {
	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/generate", genHandler)

	fmt.Print("\n✨Starting server at port 8080.✨\n\nClick here 👉: http://localhost:8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "On the wrong path!", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}

func genHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err error
	)

	// Parses the request and populates r.Form
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Take size and loops as an input, convert to int and use to create sudoku solution
	si := r.FormValue("Size")
	if si == "" {
		si = "9"
	}
	l := r.FormValue("Loops")
	if l == "" {
		l = "500"
	}

	size, err := strconv.Atoi(si)
	if err != nil {
		log.Fatal(err)
	}

	loops, err := strconv.Atoi(l)
	if err != nil {
		log.Fatal(err)
	}

	// Display inputs to user
	fmt.Fprintf(w, "Size = %d\n", size)
	fmt.Fprintf(w, "Loops = %d\n", loops)

	// Use size and loops as input
	s := solver.NewSudoku(size, loops)
	s.Start(w, false)
}
