# Sudoku Generator

Generate full sudoku grid from scratch using golang.

#

## Run

To generate a sudoku board locally, run:

```
make local
```

To generate a sudoku board through the server, run:

```
make api
```

#

## Parameters

Parameters can be changed in local/main.go:

```
	LOOPS := 5000
	SIZE := 9
	MULTIPLE_SOLUTIONS = true
```

Currently best results come from 9x9 sudoku grid (size = 9).

#

## Complete Grid

Complete grid can be found locally in the sudoku repository under 'boards'

```
Complete Sudoku Board at loop X
-------------------------
| 2 6 8 | 1 3 7 | 4 5 9 |
| 5 4 3 | 6 8 9 | 1 2 7 |
| 7 1 9 | 4 5 2 | 8 3 6 |
-------------------------
| 8 7 5 | 9 6 3 | 2 1 4 |
| 3 9 4 | 2 1 8 | 6 7 5 |
| 6 2 1 | 5 7 4 | 9 8 3 |
-------------------------
| 1 3 7 | 8 4 6 | 5 9 2 |
| 4 5 2 | 3 9 1 | 7 6 8 |
| 9 8 6 | 7 2 5 | 3 4 1 |
-------------------------
```

#

## Disclaimer:

This is a work in progress!
