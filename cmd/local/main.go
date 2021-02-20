package main

import (
	"log"
	"os"

	"github.com/jessicasomaiya/sudoku/pkg/solution"
)

var (
	LOOPS = 1000
	SIZE  = 9
)

func main() {
	s := solution.NewSudoku(SIZE, LOOPS)

	b, err := os.Create("../board")
	if err != nil {
		log.Fatal(err, " os.Create")
	}
	s.FillWholeBoard(b)
}
