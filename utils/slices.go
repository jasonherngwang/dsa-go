package utils

import (
	"reflect"
	"sort"
)

// Searching

// Find index of first element equal to a target value
func IndexOf[T comparable](slice []T, target T) int {
	for i, elem := range slice {
		if elem == target {
			return i
		}
	}
	return -1
}

// Find index of first element deeply equal to a target value
func IndexOfStructure[T any](slice []T, target T) int {
	for i, elem := range slice {
		if reflect.DeepEqual(elem, target) {
			return i
		}
	}
	return -1
}

// Find index of first element satisfying a condition function (predicate)
func IndexOfCondition[T any](slice []T, conditionFunc func(arg T) bool) int {
	for i, elem := range slice {
		if conditionFunc(elem) {
			return i
		}
	}
	return -1
}

// Find index of last element equal to a target value
// Iterates in reverse
func LastIndexOf[T comparable](slice []T, target T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}

// Inserting

// Insert element at index, in-place. Similar to JS splice().
// Duplicates, then overwrites the element at index
func InsertAtIndex[T any](slice *[]T, index int, value T) {
	if len(*slice) == index {
		*slice = append(*slice, value)
		return
	}
	*slice = append((*slice)[:index+1], (*slice)[index:]...)
	(*slice)[index] = value
}

// Same as above, but non-mutating
func InsertAtIndexNonmutating[T any](slice []T, index int, value T) []T {
	if len(slice) == index {
		return append(slice, value)
	}
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
}

// Same as above, but using copy()
// Adds a zero value at the end and shifts element over
func InsertAtIndexNonmutatingCopy[T any](slice []T, index int, value T) []T {
	if len(slice) == index {
		return append(slice, value)
	}
	var zeroValue T
	slice = append(slice, zeroValue)
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

func InsertSliceAtIndex[T any](slice []T, index int, sliceToInsert []T) []T {
	secondHalf := append(sliceToInsert, slice[index:]...)
	slice = append(slice[:index], secondHalf...)
	return slice
}

// Appending, Concatenating
// https://freshman.tech/snippets/go/concatenate-slices/
// Pre-allocate a slice and copy each one in
func ConcatMultipleSlices[T any](slices [][]T) []T {
	var totalLen int

	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, totalLen)

	var i int

	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}

// Removing/Deleting

// Remove element at index, in-place. Similar to JS splice().
// Requires dereferencing the pointer to mutate in-place
func RemoveAtIndex[T any](slice *[]T, index int) {
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}

// Same as above, but non-mutating
func RemoveAtIndexNonmutating[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

// Delete all elements
func RemoveAll[T any](slice *[]T) {
	*slice = (*slice)[:0]
}

// Filtering

// Filter using condition function (predicate)
func Filter[T any](slice []T, conditionFunc func(arg T) bool) []T {
	result := []T{}
	for _, elem := range slice {
		if conditionFunc(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// Filter in-place
func FilterInPlace[T any](slice *[]T, conditionFunc func(arg T) bool) {
	result := (*slice)[:0]
	for _, elem := range *slice {
		if conditionFunc(elem) {
			result = append(result, elem)
		}
	}
	*slice = (*slice)[:len(result)]
}

// Unique (de-duplication)
// Two pointers i and j. j stays at the first instance of any contiguous run of
// duplicates, while i advances. This function copies 1 instance of each unique
// value to the front of the array, and truncates it.

// For integers.
func Unique(slice []int) []int {
	sort.Ints(slice)
	j := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] == slice[j] {
			continue
		}
		j++
		slice[j] = slice[i]
	}
	return slice[:j+1]
}

// Reverse
func Reverse[T any](slice []T) []T {
	reversedSlice := []T{}
	for i := len(slice) - 1; i >= 0; i-- {
		reversedSlice = append(reversedSlice, slice[i])
	}
	return reversedSlice
}

// Reverse in place
// Pointers at both ends, moving to center
func ReverseInPlace[T any](slice *[]T) {
	for i, j := 0, len(*slice)-1; i < j; i, j = i+1, j-1 {
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	}
}

// Sliding Window
// Example: [1, 2, 3, 4] with window size 2 produces nested array:
// [ [1, 2], [2, 3], [3, 4] ]
func SlidingWindow(windowSize int, slice []int) [][]int {
	if len(slice) <= windowSize {
		return [][]int{slice}
	}

	result := make([][]int, 0, len(slice)-windowSize+1)

	// i, j are edges of window. Stop when j goes past end of slice.
	for i, j := 0, windowSize; j <= len(slice); i, j = i+1, j+1 {
		result = append(result, slice[i:j])
	}

	return result
}
