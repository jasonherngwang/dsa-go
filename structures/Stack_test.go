package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	t.Run("Push numbers onto stack", func(t *testing.T) {
		s := Stack[int]{Val: []int{1}}
		s.Push(2)
		s.Push(3)

		expected := Stack[int]{Val: []int{1, 2, 3}}

		assert.Equal(t, expected, s)
	})

	t.Run("Push strings onto stack", func(t *testing.T) {
		s := Stack[string]{Val: []string{"a"}}
		s.Push("b")
		s.Push("c")

		expected := Stack[string]{Val: []string{"a", "b", "c"}}

		assert.Equal(t, expected, s)
	})
}

func TestStackPop(t *testing.T) {
	t.Run("Push numbers onto stack", func(t *testing.T) {
		s := Stack[int]{Val: []int{1, 2, 3}}

		actual := s.Pop()
		assert.Equal(t, 3, actual)

		expected := Stack[int]{Val: []int{1, 2}}
		assert.Equal(t, expected, s)

	})

	t.Run("Push strings onto stack", func(t *testing.T) {
		s := Stack[string]{Val: []string{"a", "b", "c"}}

		actual := s.Pop()
		assert.Equal(t, "c", actual)

		expected := Stack[string]{Val: []string{"a", "b"}}
		assert.Equal(t, expected, s)

	})
}
