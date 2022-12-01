/*
Problem
------------------------------------------
Increasing Triplet Subsequence
https://leetcode.com/problems/increasing-triplet-subsequence/description/

Given an integer array nums, return true if there exists a triple of indices
(i, j, k) such that i < j < k and nums[i] < nums[j] < nums[k]. If no such
indices exists, return false.

Inputs: Array of integers
Outputs: Boolean

Rules, Requirements, Definitions
------------------------------------------
"Increasing Triplet Subsequence": Three values in increasing order, left to
right.
- There can be any number of values between the three.

1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1

Clarifying Questions
------------------------------------------
- Only 1 possible answer? No; there can be 4 in a row, but we return before we
  detect that.

Examples, Test Cases, Edge Cases
------------------------------------------
Input: nums = [1,2,3,4,5]
Output: true
[1,2,3,4,5]
 1 2 3

Input: nums = [2,1,5,0,4,6]
Output: true
[2,1,5,0,4,6]
       1 2 3

Input: nums = [5,4,3,2,1]
Output: false; no triplet exists

Data Structure
------------------------------------------
Array

Algorithm
------------------------------------------
Approach: Greedy
Time: O(N) for 1 traversal
Space: O(1)

We're interested in 3 values: low, mid, high.
- low is the lowest encoutered value.
  - It's ok if we reassign this to a new low AFTER registering a mid.
    The "old low" is still there, before mid, so that sequence of two values
	still exists. If we find a high, the sequence is complete.
- mid is the first value encountered that is greater than and follows low.
- high is the first value encountered that is greater than and follows mid.
  - The detection of high means that an increasing subsequence exists.
We initialize low and mid to infinity as a starting point.

Steps
- Use 2 pointers: low, mid. Initialize both to Infinity (MaxInt)
- Iterate over the array:
  - If elem <= low, make elem the new low.
  - If elem > low and elem <= mid, make elem the new mid.
  - If elem > mid, return true
- Return false
*/

package leetcode

import (
	"math"
)

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	low, mid := math.MaxInt, math.MaxInt

	for _, n := range nums {
		if n <= low {
			low = n
		} else if n <= mid {
			mid = n
		} else {
			// fmt.Printf("Triplet: %d %d %d\n", low, mid, n)
			return true
		}
	}
	return false
}
