# Testing Format for Go

After writing a solution to a problem, write tests by:
- Importing any required data structures
- Defining an `input` struct with the types of arguments to be passed
- Defining a generic `testCase` struct with test name, input type, and output type
- Creating a slice of `testCase`s
- Iterating over the slice and:
  - Executing the solution function and comparing the `actual` vs `expected`
  - Using the `testify` library or `reflect.DeepEqual`

```go
package leetcode

import (
	"github.com/jasonherngwang/dsa-go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type ListNode = structures.ListNode

	type input struct {
		head *ListNode
		n    int
	}

	type testCase struct {
		name string
		args input
		expected *ListNode
	}

	tests := []testCase{
		{
			name: "Remove middle node",
			args: input{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
				n: 2,
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := removeNthFromEnd(test.args.head, test.args.n)
			assert.Equal(t, test.expected, actual)
		})
	}
}
```