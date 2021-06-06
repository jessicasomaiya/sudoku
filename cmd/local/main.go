package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessicasomaiya/sudoku/pkg/solver"
)

var (
	LOOPS              = 100000
	SIZE               = 9
	MULTIPLE_SOLUTIONS = true
)

func main() {
	s := solver.NewSudoku(SIZE)

	dir := "board"
	b, err := os.Create(dir)
	if err != nil {
		log.Fatal(err, " os.Create")
	}

	s.Start(b, MULTIPLE_SOLUTIONS, LOOPS)

	fmt.Printf("\n✨Running locally✨\n\nGo to %s to see solutions \n", dir)
}
