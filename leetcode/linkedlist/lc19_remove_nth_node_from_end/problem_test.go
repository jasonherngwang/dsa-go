package leetcode

import (
	"testing"

	"github.com/jasonherngwang/dsa-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	type ListNode = structures.ListNode

	type input struct {
		head *ListNode
		n    int
	}

	type testCase struct {
		name     string
		input    input
		expected *ListNode
	}

	tests := []testCase{
		{
			name: "Remove middle node",
			input: input{
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
		{
			name: "Remove first node",
			input: input{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
				n: 3,
			},
			expected: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
		{
			name: "Remove last node",
			input: input{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
				n: 1,
			},
			expected: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := removeNthFromEnd(test.input.head, test.input.n)
			assert.Equal(t, test.expected, actual)
		})
	}
}
