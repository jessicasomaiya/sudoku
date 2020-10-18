package solution

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func InitSudoku(n int) *Sudoku {
	s := &Sudoku{
		n:      n,
		board:  make([]int, n*n),
		row:    make([]int, n),
		column: make([]int, n),
		square: make([]int, n),
	}
	return s
}

type Sudoku struct {
	n      int
	board  []int
	row    []int
	column []int
	square []int
}

func (s *Sudoku) FillPos(pos, v int) {
	s.board[pos] = v
}

func (s *Sudoku) validate() bool {
	var total int
	for j := 0; j < s.n; j++ {
		start := j * s.n
		for i := start; i < start+s.n; i++ {
			total += s.board[i]
		}
		if total == 45 {
			return true
		}
	}
	return false
}

func (s *Sudoku) printComplete(b io.Writer) {
	if s.nextFree() == -1 {

		fmt.Println("Full board!")
		s.PrintNinePretty(b)
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

// func (s *Sudoku) FillBoard() {
// 	for i := 0; i < len(s.board); i++ {
// 		pos := s.nextFree()
// 		// for num := 1; num <= s.n; num++ {
// 		// 	if s.legalMove(num, pos) {
// 		// 		s.board[pos] = num
// 		// 	}
// 		// 	continue
// 		// }
// 		for i := 0; i < s.n; i++ {
// 			num := rand.Intn(s.n)
// 			if s.isLegal(num+1, pos) {
// 				s.board[pos] = num + 1
// 				continue
// 			}
// 		}
// 	}
// }

// Returns random legal move
func (s *Sudoku) LegalMove(pos int) int {
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
	log.Fatal("no square found at pos", pos)
	return nil
}

// TODO: generateLookup()
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
	for k, _ := range s.board {
		s.board[k] = 0
	}
}

func (s *Sudoku) FillWholeBoard(loops int, b *http.ResponseWriter) {
	for i := 0; i <= loops; i++ {
		// if i%50 == 0 {
		// 	fmt.Fprintf(b, "Board at %d loops \n", i)
		// 	s.PrintNinePretty(b)
		// 	fmt.Fprintln(b)
		// }
		pos := s.nextFree()
		if pos == -1 && s.validate() {
			fmt.Fprintf(b, "\nComplete Sudoku Board at loop %d \n", i)
			s.PrintNinePretty(b)
			s.clearBoard()
			continue
		}
		value := s.LegalMove(pos)
		if value == -1 {
			// fmt.Fprintf(b, "no legal moves in pos %d \n", pos)
			// No legal move so move back in the tree
			// nextFree will be one before current one
			s.FillPos(pos-1, 0)
			s.FillPos(pos-2, 0)
			s.FillPos(pos-3, 0)
			s.FillPos(pos-4, 0)
			s.FillPos(pos-5, 0)
			s.FillPos(pos-6, 0)
			s.FillPos(pos-7, 0)
			s.FillPos(pos-8, 0)

			continue
		}
		s.FillPos(pos, value)
	}
	fmt.Fprintf(b, "\nFinished all loops = %d \n", loops)
	s.PrintNinePretty(b)
}

func (s *Sudoku) PrintNinePretty(b io.Writer) {
	fmt.Fprintf(b, "-------------------------\n")
	for k, v := range s.board {
		// The beginning of every line
		if k%s.n == 0 {
			fmt.Fprintf(b, "| %d ", v)
			continue
		}
		// Every third pos
		if (k+1)%(s.n/3) == 0 && (k+1)%s.n != 0 && (k+1)%(s.n*3) != 0 {
			fmt.Fprintf(b, "%d | ", v)
			continue
		}
		// At the end of every line
		if (k+1)%s.n == 0 && (k+1)%(s.n*3) != 0 {
			fmt.Fprintf(b, "%d | \n", v)
			continue
		}
		// Break line
		if (k+1)%(s.n*3) == 0 {
			fmt.Fprintf(b, "%d | \n", v)
			fmt.Fprintf(b, "-------------------------\n")
			continue
		}
		fmt.Fprintf(b, "%d ", v)
	}
}

// func (s *Sudoku) PrintNinePretty(b io.Writer) {
// 	fmt.Print(strings.Repeat("-", ((s.n+4)*2)), "\n")
// 	for k, v := range s.board {
// 		// The beginning of every line
// 		if k%s.n == 0 {
// 			fmt.Printf("| %d ", v)
// 			continue
// 		}
// 		// Every third pos
// 		if (k+1)%(s.n/3) == 0 && (k+1)%s.n != 0 && (k+1)%(s.n*3) != 0 {
// 			fmt.Printf("%d | ", v)
// 			continue
// 		}
// 		// At the end of every line
// 		if (k+1)%s.n == 0 && (k+1)%(s.n*3) != 0 {
// 			fmt.Printf("%d | \n", v)
// 			continue
// 		}
// 		// Break line
// 		if (k+1)%(s.n*3) == 0 {
// 			fmt.Printf("%d | \n", v)
// 			fmt.Print(strings.Repeat("-", (s.n+4)*2), "\n")
// 			continue
// 		}
// 		fmt.Printf("%d ", v)
// 	}
// }
