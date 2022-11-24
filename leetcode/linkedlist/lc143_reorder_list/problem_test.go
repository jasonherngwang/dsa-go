package leetcode

import (
	"github.com/jasonherngwang/dsa-go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reorderList(t *testing.T) {
	type ListNode = structures.ListNode

	t.Run("Works for an even number of nodes", func(t *testing.T) {
		input := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		}
		got := reorderList(input)
		want := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		}

		assert.Equal(t, want, got)
	})

	t.Run("Works for an odd number of nodes", func(t *testing.T) {
		input := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}
		got := reorderList(input)
		want := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		}

		assert.Equal(t, want, got)
	})

	t.Run("Failing test case", func(t *testing.T) {
		input := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}
		got := reorderList(input)
		want := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 42,
				Next: &ListNode{
					Val: 888,
					Next: &ListNode{
						Val: 9000,
						Next: &ListNode{
							Val: 1024,
						},
					},
				},
			},
		}

		assert.NotEqual(t, want, got)
	})
}

func Test_reorderListUsingArray(t *testing.T) {
	type ListNode = structures.ListNode

	t.Run("Works for an even number of nodes", func(t *testing.T) {
		input := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		}
		got := reorderListUsingArray(input)
		want := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		}

		assert.Equal(t, want, got)
	})

	t.Run("Works for an odd number of nodes", func(t *testing.T) {
		input := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}
		got := reorderListUsingArray(input)
		want := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		}

		assert.Equal(t, want, got)
	})

	t.Run("Failing test case", func(t *testing.T) {
		input := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}
		got := reorderListUsingArray(input)
		want := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 42,
				Next: &ListNode{
					Val: 888,
					Next: &ListNode{
						Val: 9000,
						Next: &ListNode{
							Val: 1024,
						},
					},
				},
			},
		}

		assert.NotEqual(t, want, got)
	})
}
