/*
Group Anagrams
https://leetcode.com/problems/group-anagrams/

Problem
------------------------------------------
Given an array of strings strs, group the anagrams together. You can return the
answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different
word or phrase, typically using all the original letters exactly once.

Inputs: Array of strings
- Lowercase English letters only.
Outputs: Nested array, where each subarray contains anagrams grouped from the
input array
- Order doesn't matter

Rules, Requirements, Definitions
------------------------------------------
- Strings can be empty
- Array will not be empty

"Anagrams" are words with the same letters
- Identical words are anagrams

Clarifying Questions
- Can there be duplicate strings? Yes.


Examples, Test Cases
------------------------------------------
Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]

Data Structure
------------------------------------------
Hash table: Key is unique letter combo, value is array of anagrams of that combo

Algorithm
------------------------------------------
Approach 1: Sort each string and add to appropriate group in hash table as we
traverse the array of strings.
Time: Assume N elements, and longest string has length K
- Sorting the longest string takes O(K log K).
- Sorting all strings takes O(N * K * log K)
Space: Hash table (object) stores N strings of up to length K. O(N*K).

Steps
- Initialize empty object to store the result.
- Iterate through array of strings:
  - Sort the characters in the string.
  - If this sorted string is a key in the result object, append the string
    (unsorted) to the array that is the property value associated with that
    object.
  - If not, create a new property:
    - Key: sorted string
    - Value: Array containing the current (unsorted) string
- Convert object into array by extracting all the values

Approach 2: Similar to Approach 1, but using char count to generate key, instead
of the sorted form of the string
Time: O(N*K). For each of N strings, we iterate over up to K chars to create the
char count key.
  - Faster than Approach 1 since we don't need to perform K log K char sorting.
Space: O(N*K). Similar hash table as Approach 1.

Helper function
- Create an alphabet array of size 26, with all initialized to 0.
  - In Go, optimize by using [26]byte instead of []string
- Iterate over the chars in the string.
  - Find the character code, e.g 98 for 'b'.
  - Subtract 97 from the code. This returns the array index for the letter.
	- In Go, find the distance between runes by subtracting them.
  - Increment the count at the array position.
- Join all values of the array into a string, using an arbitrary separator such
  as ','. This is the key for all anagrams containing the same character
  counts.
  - The separator is needed since if we have adjacent letters with the same
    digits, e.g. 1+11 vs 11+1, the keys will not be distinct.
- Return the key
- Go supports the use of any comparable type as a map key, so we can return the
  count array directly.

Steps
- Initialize an empty result array which will contain arrays of strings.
- Initialize a hash table where each key represents an anagram group, and each
  value is an array of the matching strings.
- Iterate through array of input strings
  - Generate char count key using helper method.
  - If key exists in the hash, add current string to its anagram array.
  - Else, create a new group
      - Key: char count key
      - Value: Array containing the current string
  - In Go, no need to perform this check, simply append to a non-existing array
	which will have a zero value of [].
- Extract all the groups into the result array.

Approach 3 (TOO SLOW; FAILS LEETCODE)
Create helper function to check if 2 words are anagrams.

Time: O(N^2). Comparing each string with potentially every other string. Worst
case is if each string belongs to its own group.
Space: O(N)
- Initialize result array
- Iterate through strings in input array.
  - Iterate through subarrays of result array
    - Take one string from the subarray, and check if it and the current string
      are anagrams. If so, add the current string to that subarray, and exit
      the iteration.
  - If the current string did not belong to any of the existing anagram groups,
    create a new group (new array) with the current string, and push this array
    on the result array.
    - Use a flag to determine if a group was found or not.


*/

package leetcode

import (
	"sort"
	"strings"
)

// Approach 1: Sorted chars serve as grouping key
func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	groups := make(map[string][]string)

	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		strSorted := strings.Join(chars, "")
		groups[strSorted] = append(groups[strSorted], str)
	}

	for _, group := range groups {
		result = append(result, group)
	}
	return result
}

// Approach 2: Use English alphabet letter count array as grouping key
func charCount(str string) [26]byte {
	var charCounts [26]byte
	for _, char := range str {
		charCounts[char-'a']++
	}
	return charCounts
}

func groupAnagrams2(strs []string) [][]string {
	result := [][]string{}
	groups := make(map[[26]byte][]string)

	for _, str := range strs {
		key := charCount(str)
		groups[key] = append(groups[key], str)
	}

	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
