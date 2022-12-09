/*
Given a string s, partition s such that every substring of the partition is a
palindrome.

Return all possible palindrome partitioning of s.

Input: aab

Partitions:
a|ab
aa|b
a|a|b

Output:
[
  ["aa","b"],
  ["a","a","b"]
]

Given string aab, we evaluate partitions of all size:
1: a    Remaining: ab  Add to result: a
2: aa   Remaining:  b  Add to result: aa
3: aab  NOT palindrome

If any of these are palindromes, we recursively evaluate the REMAINING
characters in the string to see if they can be partitioned into palindromes as
well.

For "1: a" above, recurse on remaining "ab"
1: a   Remaining: b  Add to result: a
2: ab  NOT palindrome

For "1: a" above, recurse on remaining "b"
1: b   Remaining:  Add to result: b

Nothing left to recurse over, so we have reached a leaf node. Add [a, a, b]

Start Index: First index of the next partition

Leaf node condition:
- When start index is the same as the input string length. AKA we have
  reached the end of the string, and are evaluating an empty string.
- We could only have reached here if all previous partitions were palindromes.

Edges: Every possible substring from the starting index onward.

Prune when: One of the partitions is not a palindrome
- We only need to evaluate the first partition (the prefix) and short circuit.
*/

package main

import "fmt"

func main() {
	isPalindrome := func(s string) bool {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-i-1] {
				return false
			}
		}
		return true
	}

	partition := func(s string) [][]string {
		result := [][]string{}

		var dfs func(startIndex int, path []string)
		dfs = func(startIndex int, path []string) {
			// Leaf node
			if startIndex == len(s) {
				fmt.Println("leaf:", path)
				result = append(result, path)
				return
			}

			// Slicing is exclusive of rightmost elem.
			for end := startIndex + 1; end <= len(s); end++ {
				prefix := s[startIndex:end]
				fmt.Printf("start %d, end %d, prefix: %s\n", startIndex, end, prefix)

				// Prune
				if !isPalindrome(prefix) {
					continue
				}

				// DFS
				path = append(path, prefix)
				fmt.Println(path)
				dfs(end, path)
				path = path[:len(path)-1]

			}
		}

		dfs(0, []string{})
		return result
	}

	fmt.Println(partition("aab"))
}
