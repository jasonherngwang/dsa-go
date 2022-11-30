/*
Problem
------------------------------------------
Sum Swap

Given two arrays of ints, swap 1 number from each so both arrays an equal sum.

Inputs: 2 arrays of integers (one or more)
Outputs: 2-element array with indexes of numbers to be swapped
[arr1Index, arr2Index]

Rules, Requirements, Definitions
------------------------------------------
- No mutation; just return the indexes.
- When we perform a swap, arr1's sum loses the difference, and arr2's sum gains
  the difference, or vice versa.
  - We need to find an element pair which has a difference of half this sum
    difference, i.e. arr1Elem - arr2Elem = (arr1Sum - arr2Sum) / 2

Clarifying Questions
------------------------------------------
- Only 1 possible answer? Yes

Examples, Test Cases, Edge Cases
------------------------------------------
Case:
- Input: [5, 3, 2, 9, 1], [1, 12, 5]
- Output: [2, 0]
  [5, 3, 1, 9, 1], [2, 12, 5]

Data Structure
------------------------------------------
Array, hash

Algorithm
------------------------------------------
Helper function: Sum of array

Approach 1: Using hash table
Time: O(N) for several passes across the arrays
Space: O(N) for extra hash

As we iterate over each element in arr1, we need to find a complement in arr2 to
satisfy the equation:
arr1Elem - arr2Elem = (arr1Sum - arr2Sum) / 2
arr2Elem = arr1Elem - (arr1Sum - arr2Sum) / 2

To look up this complement value in O(1) time, prepare a hash by first iterating
over arr2:
Key: arr1Elem = arr2Elem + (arr1Sum - arr2Sum) / 2
Value: arr2Index

To avoid integer division, double everything:
- When iterating of arr1, use 2 * arr1Elem to perform the complement lookup.
- When creating the arr2 hash, use:
  - Key: 2 * arr2Elem + (arr1Sum - arr2Sum)
  - Value: arr2Index

Approach 2: Brute force O(M*N)
- Calc sums of both arrays; find difference.
- Iterate over arr1
  - Compare each elem in arr1 with each elem in arr2; take difference
  - If difference * 2 = difference in array sums, this is the pair to swap

*/

package csg

func sumArray(s []int) int {
	sum := 0
	for _, val := range s {
		sum += val
	}
	return sum
}

func sumSwap(arr1 []int, arr2 []int) [2]int {
	sumDiff := sumArray(arr1) - sumArray(arr2)

	arr2Hash := make(map[int]int)
	for arr2Index, arr2Elem := range arr2 {
		key := 2*arr2Elem + sumDiff
		arr2Hash[key] = arr2Index
	}

	for arr1Index, arr1Elem := range arr1 {
		if arr2Index, ok := arr2Hash[2*arr1Elem]; ok {
			return [2]int{arr1Index, arr2Index}
		}
	}

	return [2]int{-1, -1}
}
