/*
Find Minimum in Rotated Sorted Array

A sorted array of integers was rotated at an unknown pivot.
For example, [10, 20, 30, 40, 50] becomes [30, 40, 50, 10, 20].
Find the index of the minimum element in this array (first index if there are
duplicates).

Input: [30, 40, 50, 10, 20]
Output: 3
Explanation: the smallest element is 10 and its index is 3.

Input: [3, 5, 7, 11, 13, 17, 19, 2]
Output: 7
Explanation: the smallest element is 2 and its index is 7.

Input: [10, 20, 30, 40, 50]
Output: 0
Explanation: There was no rotation

Notes
- If no rotation: last element is largest; first is smallest
- If rotated:
  - Last element is NOT largest. Largest is somewhere in the middle.

Search space: Ends of the array
Feasibility check:
- If elem > last elem, rotation has occurred. Discard the left half of the
  search space. Else discard the right half.

Edge case: There was no rotation; return the first element.
*/

package main

import "fmt"

// Using for l <= r
func findMinRotated(arr []int) int {
	l, r := 0, len(arr)-1
	res := -1

	for l <= r {
		m := l + (r-l)/2

		// Optimization: If we bounded inside the rotated portion, return the
		// first elem.
		if arr[l] < arr[r] {
			return l
		}

		if arr[m] <= arr[len(arr)-1] {
			// We are in the rotated portion
			res = m
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return res
}

// Using for l < r
func findMinRotated2(arr []int) int {
	l, r := 0, len(arr)-1

	for l < r {
		m := l + (r-l)/2
		if arr[m] <= arr[len(arr)-1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func main() {
	fmt.Println(findMinRotated([]int{50, 10, 20, 30, 40})) // 1
	fmt.Println(findMinRotated([]int{30, 40, 50, 10, 20})) // 3
	fmt.Println(findMinRotated([]int{30, 40, 50, 10, 10})) // 3
	fmt.Println(findMinRotated([]int{10, 20, 30, 40, 50})) // 0
	fmt.Println(findMinRotated([]int{10, 10, 30, 40, 10})) // 4

	fmt.Println(findMinRotated2([]int{50, 10, 20, 30, 40})) // 1
	fmt.Println(findMinRotated2([]int{30, 40, 50, 10, 20})) // 3
	fmt.Println(findMinRotated2([]int{30, 40, 50, 10, 10})) // 3
	fmt.Println(findMinRotated2([]int{10, 20, 30, 40, 50})) // 0
	fmt.Println(findMinRotated2([]int{10, 10, 30, 40, 10})) // 4
}
