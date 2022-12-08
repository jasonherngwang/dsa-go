/*
Generate All Letter Combinations from a Phone Number
Given a phone number that contains digits 2-9, find all possible letter
combinations the phone number could translate to.

Input: 5-6
Output: ["jm","jn","jo","km","kn","ko","lm","ln","lo"]

5: JKL
6: MNO

Result strings are length 2
startIndex: Initially 0; stop when it reaches 2

Branching
Nested branches; one for each digit in the phone number

Steps
- Create a map of phone digit to letter array.
- DFS
  - If leaf node (path is same length as phone num), add path to combos array.
  - Else:
  	- Select the next digit in the phone number whose letters we want to
      iterate over.
    - Retrieve letters and iterate over them:
      - Push on path
      - Recursive DFS using the new path
      - Pop off path
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	digitToLetter := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	letterCombinationsOfPhoneNumber := func(digits string) []string {
		combos := []string{}
		digitNums := strings.Split(digits, "")

		var dfs func(index int, path []string)
		dfs = func(index int, path []string) {
			// Leaf node
			if index == len(digits) {
				combos = append(combos, strings.Join(path, ""))
				return
			}

			nextChar := digitNums[index]
			letters := digitToLetter[nextChar]
			for _, letter := range letters {
				path = append(path, letter)
				dfs(index+1, path)
				path = path[:len(path)-1]
			}
		}

		dfs(0, []string{})
		return combos
	}

	fmt.Println(letterCombinationsOfPhoneNumber("56"))
}
