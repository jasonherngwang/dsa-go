/*
Problem
------------------------------------------
Maximum Subarray
https://leetcode.com/problems/maximum-subarray/

Given an integer array nums, find the subarray which has the largest sum and
return its sum.

"Subarray": Contiguous elements; min length 1

Inputs: Array of numbers
- Can be negative, 0, positive
Outputs: Integer

Rules, Requirements, Definitions
------------------------------------------
1 <= nums.length <= 105
-104 <= nums[i] <= 104


Examples, Test Cases, Edge Cases
------------------------------------------
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Input: nums = [1]
Output: 1

Input: nums = [5,4,-1,7,8]
Output: 23

Data Structure
------------------------------------------
Array and pointers

Algorithm
------------------------------------------
Time: O(N) for 1 traversal
Space: O(1) for pointers

Use Kadane's algo:
- If current value is greater than sum of all previous elements, then discard
  all those previous elements and start over with the current.

Steps:
- Initialize variable to track current sum and max observed sum.
- Iterate over array:
  - If elem > sum, discard previous elems.
    - Set sum to current elem.
  - Else, add current elem to current sum
  - Compare current and max sum; overwrite if larger
- Return max sum
*/

package leetcode

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	sum := nums[0]

	for _, n := range nums[1:] {
		sum = max(n, sum+n)
		maxSum = max(sum, maxSum)
	}

	return maxSum
}
