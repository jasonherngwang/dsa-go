/*
Bubble Sort
"Keep bubbling the largest elements to the right."
- Stable; an element will not bubble past equal or greater elements.
- In-place
- Time: O(N^2)
- Space: O(1)

Steps
- The sorted partiton is on the right; the unsorted on the left.
- Repeated iterate over the unsorted partition:
  - Keep swapping the element with the one on its right, as long as it is
    larger. Think of this as "bubbling" it to the front of the sorted partition.
  - If we complete an iteration with any swaps, sorting is complete.
    - Use a flag to detect this.

Optimization: Track the final destination of the element being bubbled. On the
next iteration we don't need to iterate past that point, since everything
after that is already sorted.
*/

package utils

import (
	"golang.org/x/exp/constraints"
)

func bubbleSort[T constraints.Ordered](slice []T) {
	for i := len(slice) - 1; i > 0; i-- {
		swapOccurred := false
		lastSwapIndex := 0

		for j := 0; j < i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapOccurred = true
				lastSwapIndex = j + 1
			}
		}

		if swapOccurred {
			i = lastSwapIndex // optimize by skipping iterations
		} else {
			return
		}
	}
}
