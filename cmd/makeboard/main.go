package main

import (
	"log"
	"os"

	"github.com/jessicasomaiya/sudoku/pkg/api"
	"github.com/jessicasomaiya/sudoku/pkg/solution"
)

var (
	LOOPS = 1000
	SIZE  = 9
)

func main() {

	s := solution.NewSudoku(SIZE, LOOPS)

	b, err := os.Create("board")
	if err != nil {
		log.Fatal(err, " os.Create")
	}
	s.FillWholeBoard(b)

	api.Server()

	// This isn't working!
	// runtime := flag.String("runtime", "go", "go|api runner")
	// switch *runtime {
	// case "go":
	// 	b, err := os.Create("board")
	// 	if err != nil {
	// 		log.Fatal(err, " os.Create")
	// 	}
	// 	s.FillWholeBoard(b)
	// case "api":
	// 	api.Server()
	// }
}
