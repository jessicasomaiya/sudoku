package solution

import (
	"fmt"
	"log"
	"strings"
)

func InitSudoku(n int) *Sudoku {
	s := &Sudoku{
		n:      n,
		board:  make([]int, n*n),
		row:    make([]int, n),
		column: make([]int, n),
		square: make([]int, n),
	}
	s.generateSlice(n)
	return s
}

type Sudoku struct {
	n        int
	board    []int
	row      []int
	column   []int
	square   []int
	one_to_n []int
}

func (s *Sudoku) generateSlice(n int) {
	for i := 1; i <= n; i++ {
		s.one_to_n = append(s.one_to_n, i)
	}
}

func (s *Sudoku) printComplete() {
	if s.nextFree() == -1 {
		s.printBoard()
	}
}

func (s *Sudoku) printBoard() {
	for i := 0; i < s.n; i++ {
		fmt.Println(s.board[i*s.n : (i+1)*s.n])
	}
}

func contains(slice []int, e int) bool {
	for _, v := range slice {
		if v == e {
			return true
		}
	}
	return false
}

func (s *Sudoku) nextFree() int {
	for k, v := range s.board {
		if v == 0 {
			return k
		}
	}
	return -1
}

func (s *Sudoku) FillBoard() {
	for i := 0; i < len(s.board); i++ {
		pos := s.nextFree()
		for _, num := range s.one_to_n {
			if s.legalMove(num, pos) {
				s.board[pos] = num
			}
			continue
		}
		// for i := 0; i < s.n; i++ {
		// 	num := rand.Intn(s.n)
		// 	if s.legalMove(num+1, pos) {
		// 		s.board[pos] = num + 1
		// 		continue
		// 	}
		// }
	}
}

func (s *Sudoku) posBoard() {
	for i := 0; i < len(s.board); i++ {
		s.board[i] = i
	}
}

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
	log.Fatal("no square found")
	return nil
}

// TODO: generateLookup()
func (s *Sudoku) legalMove(check, pos int) bool {
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

func (s *Sudoku) PrintNinePretty() {
	fmt.Print(strings.Repeat("-", (s.n+4)*2), "\n")
	for k, v := range s.board {
		// The beginning of every line
		if k%s.n == 0 {
			fmt.Printf("| %d ", v)
			continue
		}
		// Every third pos
		if (k+1)%(s.n/3) == 0 && (k+1)%s.n != 0 && (k+1)%(s.n*3) != 0 {
			fmt.Printf("%d | ", v)
			continue
		}
		// At the end of every line
		if (k+1)%s.n == 0 && (k+1)%(s.n*3) != 0 {
			fmt.Printf("%d | \n", v)
			continue
		}
		// Break line
		if (k+1)%(s.n*3) == 0 {
			fmt.Printf("%d | \n", v)
			fmt.Print(strings.Repeat("-", (s.n+4)*2), "\n")
			continue
		}
		fmt.Printf("%d ", v)
	}
}
