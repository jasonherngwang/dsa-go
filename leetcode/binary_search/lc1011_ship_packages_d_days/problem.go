/*
Condition function `feasible`
- Returns true if a ship with the specified capacity can ship all packages within `days` days.
- Shifts weights off the array, and increments the number of days.
- At any point, if it took too many days, return false.

Binary search
Search space: All possible ship capacities
- Min: largest weight
- Max: sum of all weights
*/

func max(slice []int) int {
	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func shipWithinDays(weights []int, days int) int {
	feasible := func(capacity int) bool {
		d := 1
		total := 0
		for _, weight := range weights {
			total += weight
			if total > capacity {
				total = weight
				d += 1
				if d > days {
					return false
				}
			}
		}
		return true
	}

	l, r := max(weights), sum(weights)
	for l < r {
		mid := l + (r-l)/2
		if feasible(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}