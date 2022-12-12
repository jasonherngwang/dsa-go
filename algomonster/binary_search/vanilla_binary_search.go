/*
Vanilla BS
Given sorted array, find target
*/
package main

import "fmt"

// Explicitly checking mid
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	// Use <= for edge case of 1-elem array were left == right initially.
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if target < arr[mid] {
			// Note: We don't always want to discard mid!
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

// Splitting until mid is the only one left
func binarySearch2(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2
		if target <= arr[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if arr[left] != target {
		return -1
	}
	return left
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(binarySearch(arr, 3))  // 2
	fmt.Println(binarySearch2(arr, 3)) // 2

	fmt.Println(binarySearch(arr, 42))  // -1
	fmt.Println(binarySearch2(arr, 42)) // -1
}
