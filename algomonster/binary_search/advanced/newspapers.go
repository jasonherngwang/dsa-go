/*
Newspapers
You have a line of newspapers, each of which takes a certain amount of time to
read. You have X coworkers to read them. How can you batch them to minimize
reading time?
- Cannot re-order the papers
- Must batch adjacent papers
- Coworkers read their assigned papers in sequence (in series)
- Multiple coworkers can work at the same time (in parallel)

Input: newspapers_read_times = [7,2,5,10,8], num_coworkers = 2
Output: 18
- Coworker 1 reads 7 + 2 + 5 = 14
- Coworker 2 reads 10 + 8 = 18
The fastest the entire job can get done is in 18 minutes.

Constraints
1 <= newspapers_read_times.length <= 10^5
1 <= newspapers_read_times[i] <= 10^5
1 <= num_coworkers <= 10^5

Search Space
- The minimum time is the largest value. This is the lower bound of the search
  space range.
- The upper bound of the range is the sum of the entire array. As if you only
  had 1 coworker who had to read them all.

Feasibilty of current search space value under consideration
Given a proposed minimum time, is it possible for the coworkers to read all the
papers in that time (or faster)?
- Initialize num coworkers required to 0.
- Iterate over the array, gathering/summing the read time into a current sum
- If adding the next elem would cause the current sum to exceed to current
  value under consideration:
  - Increment the num coworkers required
  - If the num coworkers > allowed, this search space value fails. Return false.
  - Reset the current sum to that next element which caused the overflow.

Time:
- Overall O(N log N)
  - O(N) for max() and sum() to determine search space bounds
  - O(log N) for binary search
  - O(N) for feasibility check
*/

package main

import (
	"fmt"
)

func max(ints []int) int {
	max := ints[0]
	for _, val := range ints[1:] {
		if val > max {
			max = val
		}
	}
	return max
}

func sum(ints []int) int {
	sum := 0
	for _, val := range ints {
		sum += val
	}
	return sum
}

func main() {
	newspapersSplit := func(newspapersReadTimes []int, numCoworkers int) int {
		feasible := func(minTime int) bool {
			coworkersRequired := 0
			timeTaken := 0

			for _, readTime := range newspapersReadTimes {
				if timeTaken+readTime > minTime {
					timeTaken = 0
					coworkersRequired++
				}
				timeTaken += readTime
			}
			// Leftover time; need 1 more worker
			if timeTaken > 0 {
				coworkersRequired++
			}

			return coworkersRequired <= numCoworkers
		}

		// Search space
		left := max(newspapersReadTimes)
		right := sum(newspapersReadTimes)
		minReadtime := -1

		for left <= right {
			mid := left + (right-left)/2
			fmt.Println(left, mid, right)
			if feasible(mid) {
				// fmt.Printf("feasible with %d\n", mid)
				minReadtime = mid
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		return minReadtime
	}

	fmt.Println(newspapersSplit([]int{7, 2, 5, 10, 8}, 2)) // 18
	fmt.Println(newspapersSplit([]int{2, 3, 5, 7}, 3))     // 3
}
