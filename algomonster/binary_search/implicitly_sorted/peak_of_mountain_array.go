/*
Mountain array
- 3+ elems
- Peak (largest value) at index k.
- Monotonic increase to peak, monotonic decrease afterward. No duplicates.

Find the index of the peak element. Assume there is only one peak element.

Input: 0 1 2 3 2 1 0
Output: 3
Explanation: the largest element is 3 and its index is 3.

Search space: 2nd elem to 2nd-to-last elem
Feasibility: Elem > prev and Elem > next
- If this is satisfied, return early
- Else check if the elem is on the ascent (elem > prev) or descent (elem < prev)
*/

package main

import "fmt"

func peakOfMountainArray(arr []int) int {
	left, right := 1, len(arr)-2

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] > arr[mid-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// More concise
func peakOfMountainArray2(arr []int) int {
	left, right := 1, len(arr)-2

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	fmt.Println(peakOfMountainArray([]int{0, 1, 2, 3, 2, 1, 0}))

	fmt.Println(peakOfMountainArray2([]int{0, 1, 2, 3, 2, 1, 0}))
}
