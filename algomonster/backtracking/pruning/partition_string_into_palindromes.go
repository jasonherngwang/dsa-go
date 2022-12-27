/*
Given a string s, partition s such that every substring of the partition is a
palindrome.

Return all possible palindrome partitioning of s.
- A partition cannot be empty.

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

The children of our root node are the substrings from index 0 to last index:
a, aa, and aab

To branch, set aside that first substring, and recurse on the remainder:
Reserve a, Recurse on ab => Path is ['a']
	Reserve a, Recurse on b => Path is ['a', 'a']
		Reserve b, Recurse on '' => Success; leaf node. Return ['a', 'a', 'b']

Reserve aa, Recurse on b => Path is ['aa']
	Reserve b, Recurse on '' => Success; leaf node. Return ['aa', 'b']

Reserve aab, Recurse on '' => Fail. Prune.

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
				// We increment by `end`, not always by 1, since the prefix
				// length changes
				dfs(end, path)
				path = path[:len(path)-1]

			}
		}

		dfs(0, []string{})
		return result
	}

	fmt.Println(partition("aab"))
}
