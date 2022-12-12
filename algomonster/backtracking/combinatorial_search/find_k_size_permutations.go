/*
Given and array of integers, find all permutations of size K.
- All elements will be unique.
- In a permutation, each element can be used once. No repetition.
- The order of the permutations in the result array does not matter.

"Permutation" means that order matters. For "combinations", order doesn't matter.
For array [1, 2]:
- [1, 2] and [2, 1] are distinctly different permutations.
- [1, 2] and [2, 1] are considered the same combination.

For array [1, 2, 3] and size K=3, there are 3! = 3*2*1 = 6 permutations:
[1, 2, 3]
[1, 3, 2]
[2, 1, 3]
[2, 3, 1]
[3, 1, 2]

For array [1, 2, 3] and size K=2, there are 3*2 = 6 permutations:
[1, 2]
[1, 3]
[2, 1]
[2, 3]
[3, 1]

Approach: Depth-First Search (DFS)
We use DFS to explore every possibility.

Stopping point
------------------------------------------
We must return permutations of length K, not the length of the entire array.
This helps us define a base case (stopping point) in the recursion process.
- Base case: The gathered elements so far have length K
- If we've reached the base case, add the current permutation to some global
  results array, and return without performing additional recursive calls.

If we haven't reached a stopping point
------------------------------------------
We take many different paths down the tree, each resulting in a unique
permutation. We use an array to track the "path so far". At each node, we
push the node's value onto the array and explore the branches that start with
that value. When we're done exploring, we pop that value off and explore another
branch.

Time Complexity
------------------------------------------
Time: O(N) to create and visit every node
Space: O(log N) for DFS call stack. This is the height of the tree.
*/

package main

import (
	"fmt"
)

func main() {
	findPermutations := func(nums []int, k int) [][]int {
		// "Global" variable
		result := &[][]int{}

		var dfs func(perm []int, remaining []int)
		dfs = func(perm []int, remaining []int) {
			// Base case
			if len(perm) == k {
				*result = append(*result, perm)
				return
			}

			// Every time we consider a number, the rest of the numbers are
			// treated as `remaining` since we don't want the original number
			// to appear inside its own permutations.
			for i, num := range remaining {
				// In Go, you can't pass more that 2 args to append()
				remainingCopy := append([]int{}, remaining[:i]...)
				remainingCopy = append(remainingCopy, remaining[i+1:]...)

				dfs(
					append(perm, num),
					append(remainingCopy),
				)
			}
		}

		dfs([]int{}, nums)
		return *result
	}

	fmt.Println(findPermutations([]int{1, 2, 3}, 3))
	fmt.Println(findPermutations([]int{1, 2, 3}, 2))
	fmt.Println(findPermutations([]int{1, 2, 3}, 1))
}
