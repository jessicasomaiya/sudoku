package solver

import (
	"math"

	"github.com/jessicasomaiya/sudoku/pkg/helpers"
)

func (s *Sudoku) isSquareNumber(i int) bool {
	s.rootN = math.Sqrt(float64(i))
	_, r := math.Modf(s.rootN)

	return r == 0
}

// isLegal looks at every row, column and square and checks that
//  the int 'check' is not already on that section
func (s *Sudoku) isLegal(check, pos int) bool {
	s.setRow(pos)
	s.setColumn(pos)
	s.setSquare(pos)
	if helpers.Contains(s.row, check) {
		return false
	}
	if helpers.Contains(s.column, check) {
		return false
	}
	if helpers.Contains(s.square, check) {
		return false
	}
	return true
}
