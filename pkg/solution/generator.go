package solution

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Sudoku struct {
	n      int
	rootN  float64
	sumToN int // sum from 1 to n
	board  []int
	row    []int
	column []int
	square []int
	loops  int
}

// NewSudoku creates stucture of the sudoku board
func NewSudoku(n, loops int) *Sudoku {
	s := &Sudoku{
		n:      n,
		sumToN: sumToN(n),
		board:  make([]int, n*n),
		row:    make([]int, n),
		column: make([]int, n),
		square: make([]int, n),
		loops:  loops,
	}

	return s
}

// Start begins solving the sudoku - one is true if only one board is required
func (s *Sudoku) Start(w io.Writer, one bool) {

	if !s.isSquare(s.n) {
		fmt.Fprint(w, "\nSize must be a square number\n ")
		return
	}

	s.FillBoard(w, one)
}

func (s *Sudoku) FillBoard(w io.Writer, one bool) {

	for i := 0; i <= s.loops; i++ {
		pos := s.nextFree()
		if pos == -1 && s.validate() {
			// when complete board has been written, clear and start again
			fmt.Fprintf(w, "\nComplete Sudoku Board at loop %d \n", i)
			s.printPretty(w)

			if one {
				return
			}

			s.clearBoard()
			continue
		}

		value := s.legalMove(pos)
		if pos == -1 || value == -1 {
			// No legal move so move back in the tree
			s.clearPos(pos, 8)
			continue
		}

		// If pos != -1, ie there is a space on the board to be filled, fill it with value!
		s.fillPos(pos, value)
	}
	fmt.Fprintf(w, "\nFinished all loops = %d \n ", s.loops)
	s.printPretty(w)
}

func (s *Sudoku) clearPos(pos, stepsBack int) {
	for i := 0; i < stepsBack; i++ {
		s.fillPos(pos-i, 0)
	}
}

// fillPos fills in the given position with the value v in sudoku board
func (s *Sudoku) fillPos(pos, v int) {
	s.board[pos] = v
}

func (s *Sudoku) validate() bool {
	var total int

	for j := 0; j < s.n; j++ {
		start := j * s.n
		for i := start; i < start+s.n; i++ {
			total += s.board[i]
		}

		if total == s.sumToN {
			return true
		}
	}
	return false
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

// legalMove returns random legal move (???)
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

func (s *Sudoku) setSquare(pos int) {
	// For every position in the square, find the actual value
	for k, pos := range s.findSquarePos(pos) {
		s.square[k] = s.board[pos]
	}
}

func (s *Sudoku) findSquarePos(pos int) []int {
	for _, square := range squareLookup[s.n] {
		if contains(square, pos) {
			return square
		}
	}
	log.Fatal("no square found at pos ", pos)
	return nil
}

func (s *Sudoku) isSquare(i int) bool {
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
	if contains(s.row, check) {
		return false
	}
	if contains(s.column, check) {
		return false
	}
	if contains(s.square, check) {
		return false
	}
	return true
}

func (s *Sudoku) clearBoard() {
	for k := range s.board {
		s.board[k] = 0
	}
}

func (s *Sudoku) printPretty(w io.Writer) {

	fmt.Fprintf(w, "%s\n", strings.Repeat("--", 2*(s.n)-int(s.rootN)))
	for k, v := range s.board {
		// The beginning of every line
		if k%s.n == 0 {
			fmt.Fprintf(w, "| %2d ", v)
			continue
		}
		if int(s.rootN) == 0 {
			log.Fatalf("n is zero")
		}
		// Every third pos
		if (k+1)%(s.n/int(s.rootN)) == 0 && (k+1)%s.n != 0 && (k+1)%(s.n*int(s.rootN)) != 0 {
			fmt.Fprintf(w, "%2d | ", v)
			continue
		}
		// At the end of every line
		if (k+1)%s.n == 0 && (k+1)%(s.n*int(s.rootN)) != 0 {
			fmt.Fprintf(w, "%2d | \n", v)
			continue
		}
		// Break line
		if (k+1)%(s.n*int(s.rootN)) == 0 {
			fmt.Fprintf(w, "%2d | \n", v)
			fmt.Fprintf(w, "%s\n", strings.Repeat("--", 2*(s.n)-int(s.rootN)))
			continue
		}
		fmt.Fprintf(w, "%2d ", v)
	}
}
