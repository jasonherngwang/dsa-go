package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumSwap(t *testing.T) {
	type input struct {
		nums []int
	}

	type output = bool

	type testCase struct {
		name     string
		input    input
		expected output
	}

	tests := []testCase{
		{
			name: "sequence exists",
			input: input{
				nums: []int{2, 1, 5, 0, 4, 6},
			},
			expected: true,
		},
		{
			name: "sequence doesn't exist",
			input: input{
				nums: []int{2, 1, 5, 0, 4, 1},
			},
			expected: false,
		},
		{
			name: "sequence doesn't exist",
			input: input{
				nums: []int{1, 2, 2, 1},
			},
			expected: false,
		},
		{
			name: "sequence is too short",
			input: input{
				nums: []int{1, 2},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := increasingTriplet(test.input.nums)
			assert.Equal(t, test.expected, actual)
		})
	}
}
