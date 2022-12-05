/*
Insertion Sort
"Insert each element into its place in unsorted partition."
- Stable; elements will not be inserted before equal or smaller elements.
- In-place
- Time: O(N^2)
- Space: O(1)

Steps
- The sorted partiton is on the left; the unsorted on the right.
- Initially the first is the only element in the sorted partition. On each pass
  we will insert 1 element into this partition, until the entire array is sorted.
- Iterate over the remaining elements:
  - Keep swapping the element with the one on its left, as long as it is
    smaller. Think of this as "inserting" it into the sorted partition.
*/

package utils

import "golang.org/x/exp/constraints"

func insertionSort[T constraints.Ordered](slice []T) {
	for i := 0; i < len(slice); i++ {
		// i > 0 short-circuits to prevent index out of range for i-1
		for i > 0 && slice[i] < slice[i-1] {
			slice[i], slice[i-1] = slice[i-1], slice[i] // swap with previous
			i--
		}
	}
}
