package solver

import (
	"log"

	"github.com/jessicasomaiya/sudoku/packages/api/internal/helpers"
)

// fillPos fills in the given position with the value v in sudoku board
func (s *Sudoku) fillPos(pos, v int) {
	if pos < 0 {
		pos = 0
	}
	s.board[pos] = v
}

// setRow takes a position and asigns s.Row to be the row that position is on
func (s *Sudoku) setRow(pos int) {
	// rowNumber index 0
	rowNumber := pos / s.n
	s.row = s.board[(rowNumber)*s.n : (rowNumber+1)*s.n]
}

func (s *Sudoku) setColumn(pos int) {
	columnNumber := pos % s.n
	for i := 0; i < s.n; i++ {
		s.column[i] = s.board[i*s.n+columnNumber]
	}
}

func (s *Sudoku) findSquare(pos int) []int {
	for _, square := range squareLookup[s.n] {
		if helpers.Contains(square, pos) {
			return square
		}
	}
	log.Fatal("no square found at pos ", pos)
	return nil
}

func (s *Sudoku) setSquare(pos int) {
	// For every position in the square, find the actual value
	for k, pos := range s.findSquare(pos) {
		s.square[k] = s.board[pos]
	}
}

func (s *Sudoku) clearPos(pos, stepsBack int) {
	for i := 0; i < stepsBack; i++ {
		s.fillPos(pos-i, 0)
	}
}

func (s *Sudoku) clearBoard() {
	for k := range s.board {
		s.board[k] = 0
	}
}
