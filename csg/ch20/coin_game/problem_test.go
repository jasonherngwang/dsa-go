package csg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_coingameWinner(t *testing.T) {
	type input struct {
		numCoins int
		player   string
	}

	type testCase struct {
		name     string
		input    input
		expected string
	}

	tests := []testCase{
		{
			name: "1 coin",
			input: input{
				numCoins: 1,
				player:   "you",
			},
			expected: "them",
		},
		{
			name: "2 coins",
			input: input{
				numCoins: 2,
				player:   "you",
			},
			expected: "you",
		},
		{
			name: "3 coins",
			input: input{
				numCoins: 3,
				player:   "you",
			},
			expected: "you",
		},
		{
			name: "4 coins",
			input: input{
				numCoins: 4,
				player:   "you",
			},
			expected: "them",
		},
		{
			name: "42 coins",
			input: input{
				numCoins: 42,
				player:   "you",
			},
			expected: "you",
		},
		{
			name: "43 coins",
			input: input{
				numCoins: 43,
				player:   "you",
			},
			expected: "them",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := coingameWinner(test.input.numCoins, test.input.player)
			assert.Equal(t, test.expected, actual)
		})
	}
}
