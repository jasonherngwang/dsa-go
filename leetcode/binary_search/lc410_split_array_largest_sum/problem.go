/*
Split Array Largest Sum
Split array of integers into m continuous subarrays. Minimize the largest sum
among the subarrays.

Input:
nums = [7,2,5,10,8]
m = 2

Output: 18
Possible splits:
[7]        => 7,  [2,5,10,8] => 25
[7,2]      => 9,  [5,10,8]   => 23
[7,2,5]    => 14, [10,8]     => 18 (winner)
[7,2,5,10] => 24, [8]        => 8

Approach: Binary Search
Search space: Largest sum among the subarrays
- The smallest subarray is size 1, so the lower bound of the search space is the
  largest number since that will be the largest sum among the subarrays.
- The largest subarray is the entire array, i.e. when m = 1, so the upper bound
  of the search space is the sum of the entire array.

Helper function: Given a sum value, check if the array can be split so that the
max sum of any subarray is less that or equal to that value.
- General approach: Traverse through the array, greedily gathering as many
  numbers as possible into a basket until we can't fit any more, i.e.
  exceed the provided sum value. Set it aside and continue gathering. If we
  have set aside too many (>m), this sum value is not going to work.
- Our goal is NOT to find every possible combination and compare them all.

Steps
- Initialize variables to track:
  - Sum of current subarray; initialized to 0
  - Number of subarrays formed so far; initialized to 1
- Iterate over the input array:
  - Add the current number to the current sum
  - If new sum > threshold, it's time to create a new subarray.
    - Reset the current sum to the new element, since it is now the start of the
	  new subarray.
  - Increment the number of subarrays formed.
  - If the number of subarrays > m, return false.
- Return true
*/

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func maxElem(slice []int) int {
	maxVal := math.MinInt
	for _, val := range slice {
		maxVal = max(maxVal, val)
	}
	return maxVal
}

func sumElems(slice []int) int {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	return sum
}

func splitArray(nums []int, k int) int {
	feasible := func(largestSum int) bool {
		subarrayCount := 1
		currentSum := 0

		for _, num := range nums {
			currentSum += num

			if currentSum > largestSum {
				// Start a new subarray with num as its first element.
				currentSum = num
				subarrayCount++

				// If largestSum is too small
				if subarrayCount > k {
					return false
				}
			}
		}

		return true
	}

	// Binary search: Search space is a range of largestSum
	left := maxElem(nums)
	right := sumElems(nums)

	for left < right {
		mid := left + (right-left)/2

		if feasible(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}