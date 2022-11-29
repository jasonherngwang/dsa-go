package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_invertTree(t *testing.T) {
	type input struct {
		root *TreeNode
	}

	type testCase struct {
		name     string
		input    input
		expected *TreeNode
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
			expected: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   10,
					Right: &TreeNode{Val: 7},
				},
				Right: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 3},
				},
			},
		},
		{
			name: "empty tree",
			input: input{
				root: &TreeNode{},
			},
			expected: &TreeNode{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := invertTree(test.input.root)
			assert.Equal(t, test.expected, actual)
		})
	}
}
