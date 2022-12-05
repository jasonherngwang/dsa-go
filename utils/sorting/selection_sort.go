/*
Selection Sort
"Select the smallest and append it to the sorted partition."
- Not stable; an element may get swapped to the right of its duplicate.
- In-place
- Time: O(N^2)
- Space: O(1)

Steps
- The sorted partiton is on the left; the unsorted on the right.
- Initialize 1 variable to track the index of the smallest element observed
  during the current traversal.
- Repeatedly iterate over the unsorted partition of the array:
  - Record the index of the smallest element.
  - Swap the smallest element with the first element in the unsorted partition.
    It effectively becomes the last element of the sorted partition.
*/

package utils

import "golang.org/x/exp/constraints"

func selectionSort[T constraints.Ordered](slice []T) {
	// Stop 1 before end since array will already be sorted
	for i := 0; i < len(slice); i++ {
		minIndex := i
		for j := i; j < len(slice); j++ {
			if slice[j] < slice[minIndex] {
				minIndex = j
			}
		}
		// Bring smallest to front of unsorted partition
		slice[i], slice[minIndex] = slice[minIndex], slice[i]
	}
}
