package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessicasomaiya/sudoku/pkg/solution"
)

var (
	LOOPS = 100000
	SIZE  = 9
)

func main() {
	s := solution.NewSudoku(SIZE, LOOPS)
	dir := "board"
	b, err := os.Create(dir)
	if err != nil {
		log.Fatal(err, " os.Create")
	}
	// s.FillWholeBoard(b)

	s.Start(b, false)

	fmt.Printf("\n✨Running locally✨\n\nGo to %s to see solutions \n", dir)
}
