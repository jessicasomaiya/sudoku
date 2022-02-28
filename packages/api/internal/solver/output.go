package solver

import (
	"fmt"
	"io"
	"log"
	"strings"
)

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
