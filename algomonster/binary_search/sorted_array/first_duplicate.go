/*
Find the first occurence of a target in a sorted array
*/
package main

import (
	"fmt"
)

// Using l<r, r=m, final check at end
func findFirstOccurrence(arr []int, target int) int {
	l, r := 0, len(arr)-1

	for l < r {
		m := l + (r-l)/2

		if arr[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}

	if arr[l] != target {
		return -1
	}

	return l
}

// Using l<=r, r=m-1, no final check
func findFirstOccurrence2(arr []int, target int) int {
	l, r := 0, len(arr)-1
	lastSeenIndex := -1

	for l <= r {
		m := l + (r-l)/2

		if arr[m] >= target {
			if arr[m] == target {
				lastSeenIndex = m
			}
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return lastSeenIndex
}

func main() {
	fmt.Println(findFirstOccurrence([]int{1, 3, 3, 3, 3, 6, 10, 10, 10, 100}, 3))
	fmt.Println(findFirstOccurrence([]int{2, 3, 5, 7, 11, 13, 17, 19}, 6))
	fmt.Println(findFirstOccurrence([]int{2, 3, 5, 7, 11}, 2))

	fmt.Println(findFirstOccurrence2([]int{1, 3, 3, 3, 3, 6, 10, 10, 10, 100}, 3))
	fmt.Println(findFirstOccurrence2([]int{2, 3, 5, 7, 11, 13, 17, 19}, 6))
	fmt.Println(findFirstOccurrence2([]int{2, 3, 5, 7, 11}, 2))
}
