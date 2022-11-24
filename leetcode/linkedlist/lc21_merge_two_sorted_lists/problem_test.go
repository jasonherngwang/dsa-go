// https://github.com/austingebauer/go-leetcode/blob/master/merge_two_sorted_lists_21/solution_test.go

package leetcode

import (
	s "github.com/jasonherngwang/dsa-go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *s.ListNode
		l2 *s.ListNode
	}
	tests := []struct {
		name string
		args args
		want *s.ListNode
	}{
		{
			name: "merge two sorted lists",
			args: args{
				l1: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 2,
						Next: &s.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				l2: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 3,
						Next: &s.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &s.ListNode{
				Val: 1,
				Next: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 2,
						Next: &s.ListNode{
							Val: 3,
							Next: &s.ListNode{
								Val: 4,
								Next: &s.ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mergeTwoLists(tt.args.l1, tt.args.l2))
		})
	}
}

func Test_mergeTwoListsRecursive(t *testing.T) {
	type args struct {
		l1 *s.ListNode
		l2 *s.ListNode
	}
	tests := []struct {
		name string
		args args
		want *s.ListNode
	}{
		{
			name: "merge two sorted lists",
			args: args{
				l1: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 2,
						Next: &s.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
				l2: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 3,
						Next: &s.ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &s.ListNode{
				Val: 1,
				Next: &s.ListNode{
					Val: 1,
					Next: &s.ListNode{
						Val: 2,
						Next: &s.ListNode{
							Val: 3,
							Next: &s.ListNode{
								Val: 4,
								Next: &s.ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, mergeTwoListsRecursive(tt.args.l1, tt.args.l2))
		})
	}
}
