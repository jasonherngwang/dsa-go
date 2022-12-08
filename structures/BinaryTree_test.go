package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertTree(t *testing.T) {
	t.Run("Insert numbers", func(t *testing.T) {
		tree := &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 10},
		}
		tree.Insert(3)
		tree.Insert(7)

		expected := &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:  10,
				Left: &TreeNode{Val: 7},
			},
		}

		assert.Equal(t, expected, tree)
	})
}
