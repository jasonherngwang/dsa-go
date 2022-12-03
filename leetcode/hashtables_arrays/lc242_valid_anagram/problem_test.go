package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isAnagram(t *testing.T) {
	type input struct {
		s string
		t string
	}

	type output = bool

	type testCase struct {
		name     string
		input    input
		expected output
	}

	tests := []testCase{
		{
			name: "is anagram",
			input: input{
				s: "abc",
				t: "cba",
			},
			expected: true,
		},
		{
			name: "is anagram",
			input: input{
				s: "abcddd",
				t: "cdbdad",
			},
			expected: true,
		},
		{
			name: "is not anagram",
			input: input{
				s: "abcdef",
				t: "cba",
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := isAnagram(test.input.s, test.input.t)
			assert.Equal(t, test.expected, actual)
		})
	}
}
