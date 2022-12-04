/*
Koko Eating Bananas
n piles of bananas (integers)
Number of bananas per pile is provided in an array.

Time: h hours (integer)
Eating rate: k bananas/hour (integer)

Koko eats from 1 pile per session. If she eats fast and finishes the pile, she stops.

Goal: Minimize k so Koko can eat all bananas in h hours.

Approach: Binary search
Search space: Range of value for k (eating rate)
- Min: 1 (smallest possible integer)
- Max: Size of largest pile. If Koko can eat the largest pile in 1 hr, she can eat any other pile in 1 hr as well. There's no need to eat faster since she stops when completing a pile.

Condition function
Input: k
Output: boolean

- Initialize:
  - hours elapsed to 0
- Map over array:
  - Divide num bananas by k and round up.
- Sum all the new values. If total > h, return false.
- Return true

Alternative to mapping: Add one pile at a time and terminate early if total > h
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

func mapSlice(slice []int, transform func(val int) int) []int {
	result := []int{}
	for _, v := range slice {
		result = append(result, transform(v))
	}
	return result
}

func divideRoundUp(x, y int) int {
	return int(math.Ceil(float64(x) / float64(y)))
}

func minEatingSpeed(piles []int, h int) int {
	feasible := func(k int) bool {
		// Approach 1: Mapping (slow)
		// hoursRequiredPerPile := mapSlice(piles, func(val int) int {
		//     return divideRoundUp(val, k)
		// })

		// totalHoursRequired := sumElems(hoursRequiredPerPile)
		// if totalHoursRequired > h {
		//     return false
		// }

		// return true

		// Approach 2: Adding one by one (faster)
		hoursElapsed := 0
		for _, v := range piles {
			hoursElapsed += int(math.Ceil(float64(v) / float64(k)))
			if hoursElapsed > h {
				return false
			}
		}
		return true
	}

	left, right := 1, maxElem(piles)
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