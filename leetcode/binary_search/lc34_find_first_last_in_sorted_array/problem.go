/*
Problem
------------------------------------------
Find First and Last Position of Element in Sorted Array
https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

Inputs: Array of nums in non-decreasing order (may have adjacent duplicates)
Outputs: 2-elem array [first, last] indicating the first and last index of the
target value.
- If there is only 1 occurrence, first = last.
- If there are duplicates, last > first.
- [-1, -1] if not found

Constraints:
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109

Rules, Requirements, Definitions
------------------------------------------
Must be O(log N)
- This means binary search is an option.


Examples, Test Cases, Edge Cases
------------------------------------------
Example 1:
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]

Example 2:
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

Example 3:
Input: nums = [], target = 0
Output: [-1,-1]

Data Structure
------------------------------------------
Array and pointers

Algorithm
------------------------------------------
Binary Search, Performed Twice
First find left-most occurrence. Then take the array from that point onward
and find the right-most occurrence.

Time: O(2 * log N)
Space: O(1) for pointers
*/

package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left, right := 0, len(nums)-1
	first, last := -1, -1

	// Find first occurrence
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			first = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// If target doesn't exist, return early
	if first == -1 || nums[first] != target {
		return []int{-1, -1}
	}

	// Find last occurrence
	right = len(nums) - 1 // "Reset" upper bound of search space

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			last = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return []int{first, last}
}
