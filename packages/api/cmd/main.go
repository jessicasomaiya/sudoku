package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jessicasomaiya/sudoku/packages/api/cmd/handlers"
	"github.com/jessicasomaiya/sudoku/packages/api/internal/solver"
)

var (
	LOOPS              = 100000
	SIZE               = 9
	MULTIPLE_SOLUTIONS = true

	localFlag = flag.Bool("local", false, "the local flag outputs the sudoku solvers in a local file")

	staticDir = "packages/html"
)

func main() {
	flag.Parse()

	if *localFlag {
		s := solver.NewSudoku(SIZE)
		localDir := "board"

		b, err := os.Create(localDir)
		if err != nil {
			log.Fatal(err, " os.Create")
		}

		fmt.Printf("\n✨Running locally✨\n\nGo to %s to see solutions \n", localDir)
		s.Start(b, MULTIPLE_SOLUTIONS, LOOPS)
		return
	}

	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fileServer)

	http.HandleFunc("/generate", handlers.Run)

	fmt.Print("\n✨Starting server at port 8080.✨\n\nClick here 👉: http://localhost:8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
