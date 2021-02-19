package solution

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

type Sudoku struct {
	n      int
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
		board:  make([]int, n*n),
		row:    make([]int, n),
		column: make([]int, n),
		square: make([]int, n),
		loops:  loops,
	}
	return s
}

// fillPos fills in the given position with the value v in sudoku board
func (s *Sudoku) fillPos(pos, v int) {
	s.board[pos] = v
}

// validate checks if row adds up to 45 ir 9+8+7+6...
func (s *Sudoku) validate() bool {
	var total int
	for j := 0; j < s.n; j++ {
		start := j * s.n
		for i := start; i < start+s.n; i++ {
			total += s.board[i]
		}
		// TODO: generalise for any size n
		if total == 45 {
			return true
		}
	}
	return false
}

// printComplete prints the board when it is complete
func (s *Sudoku) printComplete(w http.ResponseWriter) {
	if s.nextFree() == -1 {

		fmt.Fprint(w, "Full board!\n")
		s.printNinePretty(w)
	}
}

// printBoard prints the board at it's current state
func (s *Sudoku) printBoard() {
	for i := 0; i < s.n; i++ {
		fmt.Println(s.board[i*s.n : (i+1)*s.n])
	}
}

// contains is a helper function that checks if element e is in the slice
func contains(slice []int, e int) bool {
	for _, v := range slice {
		if v == e {
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
	for _, square := range squareLookupForNine {
		if contains(square, pos) {
			return square
		}
	}
	log.Fatal("no square found at pos", pos)
	return nil
}

// TODO: generateLookup()

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

// FillWholeBoard generates as many sudoku solutions as possible in the given loops
func (s *Sudoku) FillWholeBoard(w io.Writer) {
	for i := 0; i <= s.loops; i++ {
		// if i%50 == 0 {
		// 	fmt.Fprintf(b, "Board at %d loops \n", i)
		// 	s.printNinePretty(b)
		// 	fmt.Fprintln(b)
		// }
		pos := s.nextFree()
		if pos == -1 && s.validate() {
			fmt.Fprintf(w, "\nComplete Sudoku Board at loop %d \n", i)
			s.printNinePretty(w)
			// when complete board has been written, clear and start again
			s.clearBoard()
			continue
		}
		value := s.legalMove(pos)
		if value == -1 {
			// fmt.Fprintf(w, "no legal moves in pos %d \n", pos)
			// No legal move so move back in the tree
			// nextFree will be one before current one
			s.fillPos(pos-1, 0)
			s.fillPos(pos-2, 0)
			s.fillPos(pos-3, 0)
			s.fillPos(pos-4, 0)
			s.fillPos(pos-5, 0)
			s.fillPos(pos-6, 0)
			s.fillPos(pos-7, 0)
			s.fillPos(pos-8, 0)

			continue
		}
		s.fillPos(pos, value)
	}
	fmt.Fprintf(w, "\nFinished all loops = %d \n", s.loops)
	s.printNinePretty(w)
}

// FillOneBoard generates a sudoku solution for one board and then stops
func (s *Sudoku) FillOneBoard(w io.Writer) {
	for i := 0; i <= s.loops; i++ {
		// if i%50 == 0 {
		// 	fmt.Fprintf(w, "Board at %d loops \n", i)
		// 	s.printNinePretty(b)
		// 	fmt.Fprintln(b)
		// }
		pos := s.nextFree()
		if pos == -1 && s.validate() {
			fmt.Fprintf(w, "\nComplete Sudoku Board at loop %d \n", i)
			break
		}
		value := s.legalMove(pos)
		if value == -1 {
			// fmt.Fprintf(b, "no legal moves in pos %d \n", pos)
			// No legal move so move back in the tree
			// nextFree will be one before current one
			s.fillPos(pos-1, 0)
			s.fillPos(pos-2, 0)
			s.fillPos(pos-3, 0)
			s.fillPos(pos-4, 0)
			s.fillPos(pos-5, 0)
			s.fillPos(pos-6, 0)
			s.fillPos(pos-7, 0)
			s.fillPos(pos-8, 0)
			continue
		}
		s.fillPos(pos, value)
	}
	s.printNinePretty(w)
}

// PrintPrettyNine writes a lovely 9x9 board from what's in s.board
func (s *Sudoku) printNinePretty(w io.Writer) {
	fmt.Fprintf(w, "-------------------------\n")
	for k, v := range s.board {
		// The beginning of every line
		if k%s.n == 0 {
			fmt.Fprintf(w, "| %d ", v)
			continue
		}
		// Every third pos
		if (k+1)%(s.n/3) == 0 && (k+1)%s.n != 0 && (k+1)%(s.n*3) != 0 {
			fmt.Fprintf(w, "%d | ", v)
			continue
		}
		// At the end of every line
		if (k+1)%s.n == 0 && (k+1)%(s.n*3) != 0 {
			fmt.Fprintf(w, "%d | \n", v)
			continue
		}
		// Break line
		if (k+1)%(s.n*3) == 0 {
			fmt.Fprintf(w, "%d | \n", v)
			fmt.Fprintf(w, "-------------------------\n")
			continue
		}
		fmt.Fprintf(w, "%d ", v)
	}
}
