package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	t.Run("Enqueue numbers", func(t *testing.T) {
		q := Queue[int]{Data: []int{1}}
		q.Enqueue(2)
		q.Enqueue(3)

		expected := Queue[int]{Data: []int{1, 2, 3}}

		assert.Equal(t, expected, q)
	})

	t.Run("Enqueue strings", func(t *testing.T) {
		q := Queue[string]{Data: []string{"a"}}
		q.Enqueue("b")
		q.Enqueue("c")

		expected := Queue[string]{Data: []string{"a", "b", "c"}}

		assert.Equal(t, expected, q)
	})
}

func TestDequeue(t *testing.T) {
	t.Run("Dequeue numbers", func(t *testing.T) {
		q := Queue[int]{Data: []int{1, 2, 3}}

		actual := q.Dequeue()
		assert.Equal(t, 1, actual)

		expected := Queue[int]{Data: []int{2, 3}}
		assert.Equal(t, expected, q)
	})

	t.Run("Dequeue strings", func(t *testing.T) {
		q := Queue[string]{Data: []string{"a", "b", "c"}}

		actual := q.Dequeue()
		assert.Equal(t, "a", actual)

		expected := Queue[string]{Data: []string{"b", "c"}}
		assert.Equal(t, expected, q)
	})
}
