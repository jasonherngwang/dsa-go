package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestSubsectionSum(t *testing.T) {
	type input struct {
		nums []int
	}

	type testCase struct {
		name     string
		input    input
		expected int
	}

	tests := []testCase{
		{
			name: "array of numbers",
			input: input{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			expected: 6,
		},
		{
			name: "array of numbers",
			input: input{
				nums: []int{5, 4, -1, 7, 8},
			},
			expected: 23,
		},
		{
			name: "array of 1 number",
			input: input{
				nums: []int{1},
			},
			expected: 1,
		},
		{
			name: "array of 0",
			input: input{
				nums: []int{0},
			},
			expected: 0,
		},
		{
			name: "array of negative numbers",
			input: input{
				nums: []int{-1, -2, -3},
			},
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := maxSubArray(test.input.nums)
			assert.Equal(t, test.expected, actual)
		})
	}
}
