package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetRow(t *testing.T) {
	for _, tc := range []struct {
		name     string
		n        int
		pos      int
		board    []int
		expected []int
	}{
		{
			name: "first row",
			n:    9,
			pos:  9,
			board: []int{
				6, 5, 3, 9, 1, 7, 2, 8, 4,
				8, 1, 2, 3, 4, 6, 9, 7, 5,
				4, 7, 9, 8, 2, 5, 1, 6, 3,

				2, 9, 6, 5, 8, 4, 7, 3, 1,
				3, 8, 1, 6, 7, 9, 5, 4, 2,
				5, 4, 7, 2, 3, 1, 8, 9, 6,

				9, 6, 8, 1, 5, 3, 4, 2, 7,
				7, 2, 5, 4, 6, 8, 3, 1, 9,
				1, 3, 4, 7, 9, 2, 6, 5, 8,
			},
			expected: []int{8, 1, 2, 3, 4, 6, 9, 7, 5},
		},
		{
			name: "last row",
			n:    9,
			pos:  80,
			board: []int{
				6, 5, 3, 9, 1, 7, 2, 8, 4,
				8, 1, 2, 3, 4, 6, 9, 7, 5,
				4, 7, 9, 8, 2, 5, 1, 6, 3,

				2, 9, 6, 5, 8, 4, 7, 3, 1,
				3, 8, 1, 6, 7, 9, 5, 4, 2,
				5, 4, 7, 2, 3, 1, 8, 9, 6,

				9, 6, 8, 1, 5, 3, 4, 2, 7,
				7, 2, 5, 4, 6, 8, 3, 1, 9,
				1, 3, 4, 7, 9, 2, 6, 5, 8,
			},
			expected: []int{1, 3, 4, 7, 9, 2, 6, 5, 8},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSudoku(tc.n)
			s.board = tc.board
			s.setRow(tc.pos)

			assert.Equal(t, tc.expected, s.row)

		})
	}

}

func TestSetColumn(t *testing.T) {
	for _, tc := range []struct {
		name     string
		n        int
		pos      int
		board    []int
		expected []int
	}{
		{
			name: "first column",
			n:    9,
			pos:  9,
			board: []int{
				6, 5, 3, 9, 1, 7, 2, 8, 4,
				8, 1, 2, 3, 4, 6, 9, 7, 5,
				4, 7, 9, 8, 2, 5, 1, 6, 3,

				2, 9, 6, 5, 8, 4, 7, 3, 1,
				3, 8, 1, 6, 7, 9, 5, 4, 2,
				5, 4, 7, 2, 3, 1, 8, 9, 6,

				9, 6, 8, 1, 5, 3, 4, 2, 7,
				7, 2, 5, 4, 6, 8, 3, 1, 9,
				1, 3, 4, 7, 9, 2, 6, 5, 8,
			},
			expected: []int{6, 8, 4, 2, 3, 5, 9, 7, 1},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSudoku(tc.n)
			s.board = tc.board

			s.setColumn(tc.pos)

			assert.Equal(t, tc.expected, s.column)
		})
	}

}

func TestSetSquare(t *testing.T) {
	for _, tc := range []struct {
		name     string
		n        int
		pos      int
		board    []int
		expected []int
	}{
		{
			name: "second square",
			n:    9,
			pos:  12,
			board: []int{
				6, 5, 3, 9, 1, 7, 2, 8, 4,
				8, 1, 2, 3, 4, 6, 9, 7, 5,
				4, 7, 9, 8, 2, 5, 1, 6, 3,

				2, 9, 6, 5, 8, 4, 7, 3, 1,
				3, 8, 1, 6, 7, 9, 5, 4, 2,
				5, 4, 7, 2, 3, 1, 8, 9, 6,

				9, 6, 8, 1, 5, 3, 4, 2, 7,
				7, 2, 5, 4, 6, 8, 3, 1, 9,
				1, 3, 4, 7, 9, 2, 6, 5, 8,
			},
			expected: []int{9, 1, 7, 3, 4, 6, 8, 2, 5},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSudoku(tc.n)
			s.board = tc.board

			s.setSquare(tc.pos)

			assert.Equal(t, tc.expected, s.square)
		})
	}

}

func TestFindSquare(t *testing.T) {
	for _, tc := range []struct {
		name     string
		n        int
		pos      int
		expected []int
	}{
		{
			name:     "second square",
			n:        9,
			pos:      12,
			expected: []int{3, 4, 5, 12, 13, 14, 21, 22, 23},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSudoku(tc.n)
			assert.Equal(t, tc.expected, s.findSquare(tc.pos))
		})
	}

}
