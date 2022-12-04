/*
Minimum Number of Days to Make m Bouquets

Inputs:
- Array of integers indicating:
  - The number of day it will take to bloom
  - The relative position of the flower
- m, the number of bouquets to make
- k, the number of adjacent flowers required per bouquet
  - "Adjacent" means that only contiguous runs may be gathered into a bouquet.

Total num of flowers required = m * k

Find min number of days needed to make m bouquets.
- If length of array < m*k, there will never be enough blooms. Return -1.

Approach: Binary search
Search space: Range of values for days
- Min: m*k, since there is 1 bloom per day, so we need m*k to gather them all
- Max: The largest value in bloomDay. If we wait that long, all the flowers that
will bloom have already done so.

Condition function: Input daysToWait (int)
- Initialize flowersGathered to 0
- Initialize bouquetsMade to 0
- Iterate over array
  - If current value <= daysToWait, then this flower can be collected.
    Increment flowersGathered.
    - If flowersGathered == k, the bouquet is complete.
        - Increment bouquetsMade.
        - If bouquetsMade >= m return true
        - Reset flowersGathered to 0, so we can begin collecting for the next
		  bouquet.
  - Else, this flower hasn't bloomed yet. That means we can't add any more
    flowers to the current bouquet and have to start over.
    - Reset flowersGathered to 0
- Return false
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

func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}

	feasible := func(daysToWait int) bool {
		flowersGathered, bouquetsMade := 0, 0
		for _, daysToBloom := range bloomDay {
			if daysToBloom <= daysToWait {
				// Collect this flower
				flowersGathered++
				// Bouquet complete
				if flowersGathered == k {
					bouquetsMade++
					// Success
					if bouquetsMade >= m {
						return true
					}
					// Start gathering another bouquet
					flowersGathered = 0
				}
			} else {
				flowersGathered = 0
			}
		}
		return false
	}

	left, right := m*k, maxElem(bloomDay)
	for left < right {
		mid := left + (right-left)/2

		if feasible(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if feasible(left) {
		return left
	}
	return -1
}