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

func TestConcatMultipleSlices(t *testing.T) {
	t.Run("concat integer slices", func(t *testing.T) {
		intSlices := [][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9, 10},
		}

		actual := ConcatMultipleSlices[int](intSlices)
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		assert.Equal(t, expected, actual)
	})
}

func TestRemoveAtIndex(t *testing.T) {
	t.Run("integer slice", func(t *testing.T) {
		intSlice := &[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		RemoveAtIndex[int](intSlice, 3)
		expected := &[]int{1, 2, 3, 5, 6, 7, 8, 9, 10}
		assert.Equal(t, expected, intSlice)
	})
}

func TestReverse(t *testing.T) {
	t.Run("integer slice", func(t *testing.T) {
		intSlice := &[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		actual := Reverse[int](*intSlice)
		expected := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		assert.Equal(t, expected, actual)
	})
}

func TestReverseInPlace(t *testing.T) {
	t.Run("integer slice", func(t *testing.T) {
		intSlice := &[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		ReverseInPlace[int](intSlice)
		expected := &[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		assert.Equal(t, expected, intSlice)
	})
}

func TestInsertAtIndex(t *testing.T) {
	t.Run("integer slice", func(t *testing.T) {
		intSlice := &[]int{1, 2, 3, 4, 6, 7, 8, 9, 10}

		InsertAtIndex[int](intSlice, 4, 5)
		expected := &[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		assert.Equal(t, expected, intSlice)
	})
}

func TestInsertSliceAtIndex(t *testing.T) {
	t.Run("integer slice", func(t *testing.T) {
		intSlice := []int{1, 2, 3, 7, 8, 9, 10}
		sliceToInsert := []int{4, 5, 6}

		actual := InsertSliceAtIndex[int](intSlice, 3, sliceToInsert)
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		assert.Equal(t, expected, actual)
	})
}

func TestFilter(t *testing.T) {
	t.Run("integer slice", func(t *testing.T) {
		intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		predicate := func(arg int) bool {
			return arg%2 == 0
		}
		actual := Filter[int](intSlice, predicate)
		expected := []int{2, 4, 6, 8, 10}
		assert.Equal(t, expected, actual)
	})
}

func TestFilterInPlace(t *testing.T) {
	t.Run("integer slice", func(t *testing.T) {
		intSlice := &[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		predicate := func(arg int) bool {
			return arg%2 == 0
		}
		FilterInPlace[int](intSlice, predicate)
		expected := &[]int{2, 4, 6, 8, 10}
		assert.Equal(t, expected, intSlice)
	})
}

func TestUnique(t *testing.T) {
	t.Run("integer slice", func(t *testing.T) {
		intSlice := []int{1, 2, 2, 3, 3, 3}
		actual := Unique(intSlice)
		expected := []int{1, 2, 3}
		assert.Equal(t, expected, actual)
	})
}

func TestSlidingWindow(t *testing.T) {
	t.Run("integer slice", func(t *testing.T) {
		intSlice := []int{1, 2, 3, 4}
		actual := SlidingWindow(2, intSlice)
		expected := [][]int{
			[]int{1, 2},
			[]int{2, 3},
			[]int{3, 4},
		}
		assert.Equal(t, expected, actual)
	})
}
