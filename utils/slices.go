package utils

import "reflect"

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
