/*
Find the First True in a Sorted Boolean Array

An array of boolean values is divided into two sections; the left section
consists of all false and the right section consists of all true.

Find the First True in a Sorted Boolean Array of the right section,
i.e. the index of the first true element. If there is no true element, return -1.

Input: arr = [false, false, true, true, true]

Output: 2

Explanation: first true's index is 2.
*/
package main

import "fmt"

/*
Approach 1
- Keep satisfactory element in the search range: right = mid
- Must use loop condition left < right, or else infinite loop because right
  will never decrement, so left < right will always be true.
- Edge case: Single-element array.
  - At the end, must check if last remaining element is satisfactory.
*/

func findBoundary(arr []bool) int {
	left, right := 0, len(arr)-1

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if arr[left] {
		return left
	}

	return -1
}

/*
Approach 2
- Always discard middle element on each loop.
- Keep track of leftmost/smallest satisfactory element since it may get
  discarded during right = mid - 1.
- Use left <= right to account for 1-element arrays.
*/

func findBoundary2(arr []bool) int {
	left, right := 0, len(arr)-1
	leftmostTrueIndex := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] {
			leftmostTrueIndex = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return leftmostTrueIndex
}

func main() {
	arr := []bool{false, false, true, true, true}
	fmt.Println(findBoundary(arr))  // 2
	fmt.Println(findBoundary2(arr)) // 2

	arr2 := []bool{false, false, false, false, false}
	fmt.Println(findBoundary(arr2))  // -1
	fmt.Println(findBoundary2(arr2)) // -1
}
