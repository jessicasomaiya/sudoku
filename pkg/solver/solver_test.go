package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	for _, tc := range []struct {
		name     string
		n        int
		pos      int
		board    []int
		expected bool
	}{
		{
			name: "happy path",
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
			expected: true,
		},
		{
			name: "incomplete",
			n:    9,
			pos:  12,
			board: []int{
				6, 5, 3, 9, 1, 7, 2, 8, 4,
				8, 1, 2, 3, 4, 6, 9, 7, 0,
				4, 7, 9, 8, 2, 5, 1, 6, 3,

				2, 9, 6, 5, 8, 4, 7, 3, 1,
				3, 8, 1, 6, 7, 9, 5, 4, 2,
				5, 4, 7, 2, 3, 1, 8, 9, 6,

				9, 6, 8, 1, 5, 3, 4, 2, 7,
				7, 2, 5, 4, 6, 8, 3, 1, 9,
				1, 3, 4, 7, 9, 2, 6, 5, 8,
			},
			expected: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSudoku(tc.n)
			s.board = tc.board

			assert.Equal(t, tc.expected, s.validate())
		})
	}

}

func TestNextFree(t *testing.T) {
	for _, tc := range []struct {
		name     string
		n        int
		pos      int
		board    []int
		expected int
	}{
		{
			name: "row 3",
			n:    9,
			board: []int{
				6, 5, 3, 9, 1, 7, 2, 8, 4,
				8, 1, 2, 3, 4, 6, 9, 7, 5,
				4, 7, 9, 8, 2, 5, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
			expected: 24,
		},
		{
			name: "row 2",
			n:    9,
			board: []int{
				6, 5, 3, 9, 1, 7, 2, 8, 4,
				8, 1, 2, 3, 4, 6, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
			expected: 15,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSudoku(tc.n)
			s.board = tc.board

			assert.Equal(t, tc.expected, s.nextFree())
		})
	}

}
