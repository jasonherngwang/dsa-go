/*
Given an array of integers sorted in increasing order and a target, find the
index of the first element in the array that is larger than or equal to the
target. Assume that it is guaranteed to find a satisfying number.

Input:
arr = [1, 3, 3, 5, 8, 8, 10]
target = 2
Output: 1

Explanation: the first element larger than 2 is 3 which has index 1.

Input:
arr = [2, 3, 5, 7, 11, 13, 17, 19]
target = 6
Output: 3

Explanation: the first element larger than 6 is 7 which has index 3.

Search space: The array itself
Condition function: If the value is >= the target

Use binary search
*/
package main

import "fmt"

func feasible(val, target int) bool {
	return val >= target
}

func firstNotSmaller(arr []int, target int) int {
	l, r := 0, len(arr)-1

	for l < r {
		m := l + (r-l)/2

		if feasible(arr[m], target) {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}

func main() {
	fmt.Println(firstNotSmaller([]int{1, 3, 3, 5, 8, 8, 10}, 2))
	fmt.Println(firstNotSmaller([]int{2, 3, 5, 7, 11, 13, 17, 19}, 6))
}
