package main

import (
	"flag"
	"log"
	"os"

	"github.com/jessicasomaiya/sudoku/pkg/solution"
)

var (
	LOOPS = 1000
	SIZE  = 9
)

func main() {
	runtime := flag.String("runtime", "go", "go|api runner")

	s := solution.NewSudoku(SIZE, LOOPS)

	switch *runtime {
	case "go":
		b, err := os.Create("board")
		if err != nil {
			log.Fatal(err, " os.Create")
		}
		s.FillWholeBoard(b)
	case "api":
		// server.Server()
	}

}
