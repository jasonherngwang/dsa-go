package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDepth(t *testing.T) {
	type input struct {
		root *TreeNode
	}

	type testCase struct {
		name     string
		input    input
		expected int
	}

	tests := []testCase{
		{
			name: "binary tree",
			input: input{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   1,
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val:  10,
						Left: &TreeNode{Val: 7},
					},
				},
			},
			expected: 3,
		},
		{
			name: "depth 1",
			input: input{
				root: &TreeNode{},
			},
			expected: 1,
		},
		{
			name: "empty tree",
			input: input{
				root: nil,
			},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := maxDepth(test.input.root)
			assert.Equal(t, test.expected, actual)
		})
	}
}
