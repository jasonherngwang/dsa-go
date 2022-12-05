/*
Quick Sort
""
- Divide and conquer + Recursion
- Not stable due to swaps
- In-place
- Time: O(N log N) on avg if list is usually divided in the center
  - Worst: O(N^2) if divided at the edges
- Space: O(1)

Pointers
- Arbitrarily select pivot index, e.g. last index.
- Initialize left pointer to the start of the array.
- Initialize right pointer to pivot. It can't be pivot-1, to account for
  the case where pivot is the largest element and already in the correct place,
  so we don't want to unnecessarily swap it.

Loop: Move pointers inward until them meet
- Advance left pointer until a value >= pivot.
- Decrement right pointer until a value < pivot. It may not move if pivot is
  already the largest.
- If pointers have not met, swap the elements.
  - If pointers meet, exit the loop.

Divide
- Swap pivot and meeting point.
- Now our array looks like: (smaller than pivot), pivot, (larger than pivot)
  - The elements don't have to be in order yet.
- Repeat the process (recursive call) for both sides of the pivot (not including
  the pivot itself).
*/

package utils

import (
	"golang.org/x/exp/constraints"
)

func quickSort[T constraints.Ordered](slice []T) {
	partition(slice, 0, len(slice)-1)
}

func partition[T constraints.Ordered](slice []T, start int, end int) {
	// Single element is already sorted.
	if end-start < 1 {
		return
	}

	// Arbitrarily select last element as pivot.
	pivot := end
	left, right := start, end

	for {
		// Advance pointers until they meet
		for slice[left] < slice[pivot] && left < right {
			left++
		}
		for slice[right] >= slice[pivot] && left < right {
			right--
		}
		if left == right {
			break
		}

		// Swap elements so larger-than-pivot element go the right, and
		// smaller-than-pivot element go the left.
		slice[left], slice[right] = slice[right], slice[left]
	}

	// Swap meeting point with pivot element
	// Possible to swap pivot with itself
	slice[left], slice[pivot] = slice[pivot], slice[left]

	// Recursive call on the two halves. Exclude the pivot itself.
	partition(slice, start, left-1)
	partition(slice, left+1, end)
}
