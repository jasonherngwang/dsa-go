package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert" // Or your library of choice
)

func Test_searchRange(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}

	type output = []int

	type testCase struct {
		name     string
		input    input
		expected output
	}

	tests := []testCase{
		{
			name: "No duplicates",
			input: input{
				nums:   []int{5, 7, 7, 8, 10},
				target: 8,
			},
			expected: []int{3, 3},
		},
		{
			name: "Single elem",
			input: input{
				nums:   []int{5},
				target: 5,
			},
			expected: []int{0, 0},
		},
		{
			name: "Duplicates",
			input: input{
				nums:   []int{5, 7, 7, 8, 8, 8, 10},
				target: 8,
			},
			expected: []int{3, 5},
		},
		{
			name: "Target not found",
			input: input{
				nums:   []int{5, 7, 7, 10},
				target: 8,
			},
			expected: []int{-1, -1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := searchRange(test.input.nums, test.input.target)
			assert.Equal(t, test.expected, actual)
		})
	}
}
