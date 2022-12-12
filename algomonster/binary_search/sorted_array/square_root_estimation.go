/*
Find the square root of an integer, rounded down.

Use binary search to find the first integer whose square is greater than the
target.

Search space: 0 to the target itself (to account for 1^2 = 1)
Feasibility: elem * elem > target
- At the end, return the result - 1

Test case:
- Input: 10
- Output: 3.? (desired answer is 3 after truncating decimal)
Search space: 1 to 10
Iteration 1:
- Mid = 5; 5*5 = 25 (too large; >=10)
- Right = Mid = 5
- Left = 1
Iteration 2:
- Mid = 3; 3*3 = 9 (too small; <10)
- Left = Mid + 1 = 4
- Right = 5
Iteration 3:
- Mid = 4; 4*4 = 16 (too large)
- Right = Mid = 4
- Left = 4
- Break out of loop, since Left < Right is no longer true
Return 4-1 = 3.
*/
package main

import "fmt"

func squareRoot(n int) int {
	if n == 0 {
		return -1
	}
	if n == 1 {
		return 1
	}

	l, r := 0, n

	for l < r {
		m := l + (r-l)/2

		if m*m > n {
			r = m
		} else {
			l = m + 1
		}
	}

	return l - 1
}

// Explicitly check if square equals n
func squareRoot2(n int) int {
	if n == 0 {
		return -1
	}

	l, r := 0, n
	res := -1

	for l <= r {
		m := l + (r-l)/2

		if m*m == n {
			return m
		} else if m*m > n {
			res = m // Keep track of value whose square is larger than target
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return res - 1
}

func main() {
	fmt.Println(squareRoot(16))
	fmt.Println(squareRoot(10))
	fmt.Println(squareRoot(8))
	fmt.Println(squareRoot(1))
	fmt.Println(squareRoot(0))

	fmt.Println(squareRoot2(16))
	fmt.Println(squareRoot2(10))
	fmt.Println(squareRoot2(8))
	fmt.Println(squareRoot2(1))
	fmt.Println(squareRoot2(0))
}
