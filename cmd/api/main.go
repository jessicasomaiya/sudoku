package main

import (
	"github.com/jessicasomaiya/sudoku/pkg/server"
)

var (
	LOOPS = 10000
	SIZE  = 9
)

func main() {
	server.New()
}
