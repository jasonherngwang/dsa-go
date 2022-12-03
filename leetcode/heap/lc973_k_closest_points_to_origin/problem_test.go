package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kClosest(t *testing.T) {
	type input struct {
		points [][]int
		k      int
	}

	type output = [][]int

	type testCase struct {
		name     string
		input    input
		expected output
	}

	tests := []testCase{
		{
			name: "4 points",
			input: input{
				points: [][]int{[]int{1, 3}, []int{-2, 2}, []int{4, 4}, []int{1, 1}},
				k:      2,
			},
			expected: [][]int{[]int{1, 1}, []int{-2, 2}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := kClosest(test.input.points, test.input.k)
			assert.Equal(t, test.expected, actual)
		})
	}
}
