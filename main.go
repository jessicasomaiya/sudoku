package main

func main() {
	s := solution.InitSudoku(9)
	s.fillBoard()

	s.printNinePretty()
}
