package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert" // Or your library of choice
)

var isBadVersion func(version int) bool

func Test_firstBadVersion(t *testing.T) {
	type input struct {
		n          int
		badVersion int
	}

	createBadVersionFunc := func(bad int) func(version int) bool {
		return func(version int) bool {
			return version >= bad
		}
	}

	type output = int

	type testCase struct {
		name     string
		input    input
		expected output
	}

	tests := []testCase{
		{
			name: "Array of 5 versions",
			input: input{
				n:          5,
				badVersion: 4,
			},
			expected: 4,
		},
		{
			name: "Array of 1 versions",
			input: input{
				n:          5,
				badVersion: 1,
			},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isBadVersion = createBadVersionFunc(test.input.badVersion)
			actual := firstBadVersion(test.input.n)
			assert.Equal(t, test.expected, actual)
		})
	}
}
