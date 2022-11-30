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
Space: O(1) for some extra variables

Approach 1: Kadane's algo, Greedy
- If current value is greater than sum of all previous elements, then discard
  all those previous elements and start over with the current.

Steps:
- Initialize variables to track current sum and max observed sum.
- Iterate over array:
  - If elem > sum, discard previous elems.
    - Set sum to current elem.
  - Else, add current elem to current sum
  - Compare current and max sum; overwrite if larger
- Return max sum

Approach 2: Dynamic Programming
Time: O(N) for 1 traversal
Space: O(N) since we use an additional array to store sums

Steps:
- Initialize variable to track max observed sum.
- Initialize array containing the first element from the input array. Each slot
  will hold one of these, whichever is larger:
  - The sum of the previous element and the value at the corresponding index in
    the input array
  - Just the value at the corresponding index in the input array
- Iterate over the input array, filling in the new array.
- Return max observed sum.
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

func maxSubArrayDP(nums []int) int {
	maxSum := nums[0]
	sums := []int{nums[0]}

	// i=0 means index 1 in sums; need to add 1 when accessing sums
	for i, n := range nums[1:] {
		sums = append(sums, max(n, sums[i]+n))
		maxSum = max(sums[i+1], maxSum)
	}

	return maxSum
}
