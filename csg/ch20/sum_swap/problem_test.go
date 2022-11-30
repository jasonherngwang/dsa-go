package csg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumSwap(t *testing.T) {
	type input struct {
		arr1 []int
		arr2 []int
	}

	type testCase struct {
		name     string
		input    input
		expected [2]int
	}

	tests := []testCase{
		{
			name: "arrays",
			input: input{
				arr1: []int{5, 3, 2, 9, 1},
				arr2: []int{1, 12, 5},
			},
			expected: [2]int{2, 0},
		},
		{
			name: "arrays",
			input: input{
				arr1: []int{5, 3, 3, 7},
				arr2: []int{4, 1, 1, 6},
			},
			expected: [2]int{3, 0},
		},
		{
			name: "arrays",
			input: input{
				arr1: []int{1, 2, 3, 4, 5},
				arr2: []int{6, 7, 8},
			},
			expected: [2]int{2, 0},
		},
		{
			name: "arrays",
			input: input{
				arr1: []int{10, 15, 20},
				arr2: []int{5, 30},
			},
			expected: [2]int{0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := sumSwap(test.input.arr1, test.input.arr2)
			assert.Equal(t, test.expected, actual)
		})
	}
}
