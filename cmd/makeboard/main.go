package main

import (
	"log"
	"os"

	"github.com/jessicasomaiya/sudoku/pkg/solution"
)

func main() {
	LOOPS := 500
	SIZE := 9
	s := solution.InitSudoku(SIZE)

	b, err := os.Create("boards")
	if err != nil {
		log.Fatal("output cannot be created")
	}
	// s.FillWholeBoard(LOOPS, b)
	s.FillOneBoard(LOOPS, b)
	s.HandleRequests()
}
