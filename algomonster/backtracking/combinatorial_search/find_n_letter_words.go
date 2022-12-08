/*
Find all n-letter words composed by 'a' and 'b', return them in a list of
strings in lexicographical order.

Input: 2
Output: ["aa", "ab", "ba", "bb"]

Input: 4
Output: ["aaaa", "aaab", "aaba", "aabb", "abaa", "abab", "abba", "abbb", "baaa",
"baab", "baba", "babb", "bbaa", "bbab", "bbba", "bbbb"]

Draw the tree
All permutations:
        ''
   a         b
aa   ab   ba   bb

We reach a solution when: We reach level n.

We branch when: At every node, we create 2 branches, since we care about "a" and "b".

Steps
Base case: Reach word of length n.
- Use startIndex to represent n, and increment by 1 for each level.
- "Report" by adding to global result array

Else:
- Add "a" to path array
- DFS
- Pop edge
- Repeat for "b"

Time: O(N) to create and visit every node
Space: O(log N) for DFS call stack
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	letterCombination := func(n int) []string {
		result := []string{}

		var dfs func(startIndex int, path []string)
		dfs = func(startIndex int, path []string) {
			// If solution found
			if startIndex == n {
				// Report solution
				result = append(result, strings.Join(path, ""))
				return
			}

			// For candidate in list of candidates
			for _, letter := range []string{"a", "b"} {
				// Place candidate
				path = append(path, letter)
				// Evaluate candidate at next level of tree
				dfs(startIndex+1, path)
				// Remove candidate and backtrack
				path = path[:len(path)-1]
			}
		}

		dfs(0, []string{})
		return result
	}

	fmt.Println(letterCombination(2))
}
