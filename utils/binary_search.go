/*
Binary Search
https://leetcode.com/discuss/general-discussion/786126/Python-Powerful-Ultimate-Binary-Search-Template.-Solved-many-problems

Assume input array is sorted in ascending order.

Steps
 1. Split search space into two halves.
 2. Only keep the half that likely contains the target. This becomes the new
    search space
 3. Continue splitting the search space in half until the target is found.

Linear search: O(N)
Binary search: O(log N)
*/

package utils

// General binary search
func binarySearch(slice []int) int {
	condition := func(val int) bool {
		return true
	}

	// Initialize pointers to the ends of the array.
	left, right := 0, len(slice)-1

	for left < right {
		mid := left + (right-left)/2 // Integer division

		if condition(mid) {
			// If mid itself satisfies the condition, left will keep advancing
			// until left == right. This adds a few loops.
			right = mid
		} else {
			// left skips mid because mid is already rejected.
			left = mid + 1
		}
	}

	// One or more elements could satisfy the condition.
	// left is the minimal/smallest k such that condition(k) is true.
	// We may want left-1, the other side of the boundary.
	return left
}
