package main

import (
	"github.com/jessicasomaiya/sudoku/pkg/solution"
)

func main() {
	s := solution.InitSudoku(9)
	s.FillBoard()

	s.PrintNinePretty()
}
