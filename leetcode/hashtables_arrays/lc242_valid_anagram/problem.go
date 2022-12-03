/*
Problem
------------------------------------------
Valid Anagram
https://leetcode.com/problems/valid-anagram/

Given two strings s and t, return true if t is an anagram of s, and false
otherwise.

Inputs: 2 strings
Outputs: Boolean

Rules, Requirements, Definitions
------------------------------------------
An Anagram is a word or phrase formed by rearranging the letters of a different
word or phrase, typically using all the original letters exactly once.
- Must be same length

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.

Examples, Test Cases, Edge Cases
------------------------------------------
Input: s = "anagram", t = "nagaram"
Output: true

Input: s = "rat", t = "car"
Output: false

Data Structure
------------------------------------------
Array, Hash

Algorithm
------------------------------------------
Approach 1: Letter tally
M is length of string `s`; N is length of string `t`
Time: O(N + M) for several traversals
Space: O(M) for hash

Steps
- If lengths different return false.
- Create letter tally of `s`.
- Iterate over letters in `t`.
  - If letter not in tally, return false.
  - Decrement value in tally.
    - If value is 0 delete the entry
	- If value < 0 return false
- If hash length 0, return true
- Return false

Approach 2: Optimization of Approach 1
- Use a 26-length array for the English alphabet.
- Iterate over both strings at the same time, using the same index.
  - Increment and decrement the letter count at the same time.
- Iterate over the letter count, checking if all zeroes.
*/

package leetcode

func letterTally(s string) map[rune]int {
	tally := make(map[rune]int)
	for _, letter := range s {
		tally[letter]++
	}
	return tally
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	tally := letterTally(s)

	for _, letter := range t {
		count, ok := tally[letter]
		if !ok {
			return false
		}

		if count == 1 {
			delete(tally, letter)
		} else {
			tally[letter]--
		}
	}

	if len(tally) == 0 {
		return true
	}

	return false
}

func isAnagramOptimized(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var counts [26]int

	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for _, val := range counts {
		if val != 0 {
			return false
		}
	}

	return true
}
