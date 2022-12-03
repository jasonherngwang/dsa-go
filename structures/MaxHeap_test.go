package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertMaxHeap(t *testing.T) {
	t.Run("Insert integer", func(t *testing.T) {
		m := &MaxHeap{}
		arr := []int{10, 20, 30, 1, 3, 5, 7, 9}
		for _, n := range arr {
			m.Insert(n)
		}

		expected := &MaxHeap{
			array: []int{30, 10, 20, 9, 3, 5, 7, 1},
		}

		assert.Equal(t, expected, m)
	})
}

func TestExtractMaxHeap(t *testing.T) {
	t.Run("Extract integer", func(t *testing.T) {
		m := &MaxHeap{
			array: []int{10, 20, 30, 1, 3, 5, 7, 9},
		}

		for i := 0; i < 5; i++ {
			m.Extract()
		}

		expected := &MaxHeap{
			array: []int{5, 3, 1},
		}

		assert.Equal(t, expected, m)
	})
}
