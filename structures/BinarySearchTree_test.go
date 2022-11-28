package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertBST(t *testing.T) {
	t.Run("Insert numbers", func(t *testing.T) {
		bst := &TreeNode[int]{
			Key:   5,
			Left:  &TreeNode[int]{Key: 1},
			Right: &TreeNode[int]{Key: 10},
		}
		bst.Insert(3)
		bst.Insert(7)

		expected := &TreeNode[int]{
			Key: 5,
			Left: &TreeNode[int]{
				Key:   1,
				Right: &TreeNode[int]{Key: 3},
			},
			Right: &TreeNode[int]{
				Key:  10,
				Left: &TreeNode[int]{Key: 7},
			},
		}

		assert.Equal(t, expected, bst)
	})

	t.Run("Insert strings", func(t *testing.T) {
		bst := &TreeNode[string]{
			Key:   "llama",
			Left:  &TreeNode[string]{Key: "cat"},
			Right: &TreeNode[string]{Key: "zebra"},
		}
		bst.Insert("giraffe")
		bst.Insert("platypus")

		expected := &TreeNode[string]{
			Key: "llama",
			Left: &TreeNode[string]{
				Key:   "cat",
				Right: &TreeNode[string]{Key: "giraffe"},
			},
			Right: &TreeNode[string]{
				Key:  "zebra",
				Left: &TreeNode[string]{Key: "platypus"},
			},
		}

		assert.Equal(t, expected, bst)
	})
}

func TestSearchBST(t *testing.T) {
	t.Run("Search for number", func(t *testing.T) {
		bst := &TreeNode[int]{
			Key: 5,
			Left: &TreeNode[int]{
				Key:   1,
				Right: &TreeNode[int]{Key: 3},
			},
			Right: &TreeNode[int]{
				Key:  10,
				Left: &TreeNode[int]{Key: 7},
			},
		}

		actual := true
		expected := bst.Search(7)

		assert.Equal(t, expected, actual)
	})
}
