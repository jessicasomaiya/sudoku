package main

import (
	"log"
	"os"

	"github.com/jessicasomaiya/sudoku/pkg/solution"
)

func main() {
	b, err := os.Create("boards")
	if err != nil {
		log.Fatal("output cannot be created")
	}
	s := solution.InitSudoku(9)

	// root := solution.CreateRoot(2)
	// mcts := solution.NewMCTS(root)
	// mcts.Gamble(1000)
	s.FillWholeBoard(5000, b)

	// s.PrintNinePretty(b)
}
