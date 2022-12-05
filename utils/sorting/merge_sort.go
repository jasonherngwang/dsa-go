/*
Merge Sort
"Recursively divide the array into halves. Once they are singles, merge them
back together."
- Divide and conquer + Recursion
- Stable
- Not in-place; uses additional arrays
- Time: O(N log N)
- Space: O(N) for the separate result array

Steps
Divide Step
- Divide the array in half, and recursively call merge sort on the halves.
  - Use a right-of-center middle index since slicing excludes right (in Go).
  - If there are an odd number of elements, it's OK for one "half" to only have
    one element.
  - We keep dividing until each half is just one element; then we start merging.
  - A 1-element array is considered "sorted".
- Perform one final merge of the two halves.

Merge Step
- Initialize an empty array to hold the sorted elements.
- Use 2 pointers, one at the beginning of each array.
  - If one pointer has already reached the end of its array (or the array is
	empty), we can directly	add the rest of the other array to the result, and
	return.
  - Add the smaller element to the result array, and advance the pointer.

Once the merge is complete, this merge sort recursive call is complete and
returns. Then we process the recursive call 1 level up.
*/

package utils

import (
	"golang.org/x/exp/constraints"
)

func mergeArrays[T constraints.Ordered](left, right []T) []T {
	pLeft, pRight := 0, 0
	result := make([]T, 0, len(left)+len(right))

	// Merge until both pointers have traversed their entire arrays.
	for pLeft < len(left) || pRight < len(right) {
		// If one array is empty, return the rest of the other array.
		// Else find the smaller element, add it, and advance its pointer.
		if pLeft == len(left) {
			return append(result, right[pRight:]...)
		} else if pRight == len(right) {
			return append(result, left[pLeft:]...)
		} else if left[pLeft] <= right[pRight] {
			result = append(result, left[pLeft])
			pLeft++
		} else {
			result = append(result, right[pRight])
			pRight++
		}
	}

	return result
}

func mergeSort[T constraints.Ordered](slice []T) []T {
	// Guard clause
	if len(slice) <= 1 {
		return slice
	}

	// Divide
	mid := len(slice) / 2
	left := mergeSort(slice[:mid])
	right := mergeSort(slice[mid:])

	// Merge
	return mergeArrays(left, right)
}
