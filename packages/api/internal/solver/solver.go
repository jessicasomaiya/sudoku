package solver

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/jessicasomaiya/sudoku/packages/api/internal/helpers"
)

type Sudoku struct {
	n      int
	rootN  float64
	sumToN int // sum from 1 to n
	board  []int
	row    []int
	column []int
	square []int
}

// NewSudoku creates stucture of the sudoku board
func NewSudoku(n int) *Sudoku {
	s := &Sudoku{
		n:      n,
		sumToN: helpers.SumToN(n),
		board:  make([]int, n*n),
		row:    make([]int, n),
		column: make([]int, n),
		square: make([]int, n),
	}

	return s
}

// Start begins solving the sudoku - many is false if only one board is required
func (s *Sudoku) Start(w io.Writer, many bool, loops int) {

	if !s.isSquareNumber(s.n) {
		fmt.Fprint(w, "\nSize must be a square number\n ")
		return
	}

	if _, ok := squareLookup[s.n]; !ok {
		fmt.Fprint(w, "\nSize not yet supported\n ")
		return
	}

	s.Run(w, many, loops)
}

func (s *Sudoku) Run(w io.Writer, many bool, loops int) {

	for i := 0; i <= loops; i++ {
		pos := s.nextFree()
		// when complete board has been written, clear and start again
		if pos == -1 && s.validate() {
			fmt.Fprintf(w, "\nComplete Sudoku Board at loop %d \n", i)
			s.printPretty(w)

			if !many {
				return
			}
			s.clearBoard()
			continue
		}

		// find next value to populate board
		value := s.legalMove(pos)
		if pos == -1 || value == -1 {
			// No legal move so move back in the tree
			s.clearPos(pos, s.n)
			continue
		}

		// If pos != -1, ie there is a space on the board to be filled, fill it with value!
		s.fillPos(pos, value)
	}
	fmt.Fprintf(w, "\nFinished all loops = %d \n ", loops)
	s.printPretty(w)
}

// for every row, check that
func (s *Sudoku) validate() bool {
	var total int

	for j := 0; j < s.n; j++ {
		start := j * s.n
		for i := start; i < start+s.n; i++ {
			total += s.board[i]
		}

		if total != s.sumToN {
			return false
		}

		total = 0

	}
	return true
}

// nextFree returns the next position on the board to be filled in
func (s *Sudoku) nextFree() int {
	for k, v := range s.board {
		if v == 0 {
			return k
		}
	}
	return -1
}

// legalMove returns random legal move
func (s *Sudoku) legalMove(pos int) int {
	var legal []int
	for v := 1; v <= s.n; v++ {
		if s.isLegal(v, pos) {
			legal = append(legal, v)
		}
	}
	if legal == nil {
		return -1
	}

	rand.Seed(time.Now().UnixNano())
	return legal[rand.Intn(len(legal))]
}
