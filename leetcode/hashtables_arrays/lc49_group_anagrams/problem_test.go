// https://github.com/austingebauer/go-leetcode/blob/master/group_anagrams_49/solution_test.go

package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_groupAnagrams(t *testing.T) {
	type input struct {
		strs []string
	}

	type testCase struct {
		name     string
		input    input
		expected [][]string
	}

	tests := []testCase{
		{
			name: "Group anagrams",
			input: input{
				strs: []string{
					"eat", "tea", "tan", "ate", "nat", "bat",
				},
			},
			expected: [][]string{
				{
					"bat",
				},
				{
					"eat", "tea", "ate",
				},
				{
					"tan", "nat",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := groupAnagrams(test.input.strs)
			// Result order doesn't matter; sort by 1st element for this case
			sort.Slice(actual, func(i int, j int) bool {
				return actual[i][0] < actual[j][0]
			})
			assert.Equal(t, test.expected, actual)
		})
	}
}
