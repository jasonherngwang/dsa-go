/*
Problem
------------------------------------------
Binary Search
https://leetcode.com/problems/binary-search/description/

Inputs: Array of number, Target number
Outputs: Index of target in array. If not found, return -1

Rules, Requirements, Definitions
------------------------------------------
Constraints:
1 <= nums.length <= 104
-104 < nums[i], target < 104
All the integers in nums are unique.
nums is sorted in ascending order.


Examples, Test Cases, Edge Cases
------------------------------------------
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1

Data Structure
------------------------------------------
Array and pointers

Algorithm
------------------------------------------
Time: O(log N) for binary search
Space: O(1) for pointers

Steps
- Initialize pointers left and right, to first and last index. This covers the
  entire search space.
- While left < right:
  - Calculate mid index using integer division.
  - If value at mid index is >= target, then target lies in the first half.
    Halve the search space by setting right to mid.
  - Else, target lies in the second half. Set left to mid+1.
- When left == right, the loop exits.
- If the value at left is not the target, return -1.
- Else return the value at left.

*/

package leetcode

func binarySearch(nums []int, target int) int {
	condition := func(index int) bool {
		return target <= nums[index]
	}

	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if condition(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if nums[left] != target {
		return -1
	}
	return left
}
