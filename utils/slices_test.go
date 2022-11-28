package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		intSlice := []int{1, 2, 3}

		actual := IndexOf[int](intSlice, 3)
		expected := 2
		assert.Equal(t, expected, actual)

		actual = IndexOf[int](intSlice, 4)
		expected = -1
		assert.Equal(t, expected, actual)
	})

	t.Run("strings", func(t *testing.T) {
		strSlice := []string{"1", "2", "3"}

		actual := IndexOf[string](strSlice, "3")
		expected := 2
		assert.Equal(t, expected, actual)

		actual = IndexOf[string](strSlice, "42")
		expected = -1
		assert.Equal(t, expected, actual)
	})
}

func TestIndexOfStructure(t *testing.T) {
	t.Run("slices", func(t *testing.T) {
		intSlices := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		actual := IndexOfStructure[[]int](intSlices, []int{4, 5, 6})
		expected := 1
		assert.Equal(t, expected, actual)

		actual = IndexOfStructure[[]int](intSlices, []int{42, 5, 6})
		expected = -1
		assert.Equal(t, expected, actual)
	})

	t.Run("maps", func(t *testing.T) {
		mapSlices := []map[string]int{
			{"num": 1},
			{"num": 2},
			{"num": 3},
		}

		actual := IndexOfStructure[map[string]int](mapSlices, map[string]int{"num": 2})
		expected := 1
		assert.Equal(t, expected, actual)

		actual = IndexOfStructure[map[string]int](mapSlices, map[string]int{"num": 4})
		expected = -1
		assert.Equal(t, expected, actual)
	})
}

func TestIndexOfCondition(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		intSlice := []int{1, 2, 3}

		actual := IndexOfCondition[int](intSlice, func(arg int) bool { return arg == 3 })
		expected := 2
		assert.Equal(t, expected, actual)

		actual = IndexOfCondition[int](intSlice, func(arg int) bool { return arg == 42 })
		expected = -1
		assert.Equal(t, expected, actual)
	})

	t.Run("strings", func(t *testing.T) {
		strSlice := []string{"1", "2", "3"}

		actual := IndexOfCondition[string](strSlice, func(arg string) bool { return arg == "3" })
		expected := 2
		assert.Equal(t, expected, actual)

		actual = IndexOfCondition[string](strSlice, func(arg string) bool { return arg == "42" })
		expected = -1
		assert.Equal(t, expected, actual)
	})
}

func TestRemoveAtIndex(t *testing.T) {
	t.Run("integer slice", func(t *testing.T) {
		intSlice := &[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		RemoveAtIndex[int](intSlice, 3)
		expected := &[]int{1, 2, 3, 5, 6, 7, 8, 9, 10}
		assert.Equal(t, intSlice, expected)
	})
}
