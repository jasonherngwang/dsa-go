package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Include desired assertions in exports
var Equal = assert.Equal

type TestCase[Input any] struct {
	Name     string
	Args     Input
	Expected any
}

func Run_Tests[Input any](t *testing.T, tests []TestCase[Input], testFunc func(TestCase[Input])) {
	for _, test := range tests {
		t.Run(test.Name, func(*testing.T) { testFunc(test) })
	}
}

func Test_KthLargestInitial(t *testing.T) {
	type input struct {
		k    int
		nums []int
	}

	tests := []TestCase[input]{
		{
			Name: "Initial numbers and k",
			Args: input{
				k:    3,
				nums: []int{4, 5, 8, 2},
			},
			Expected: KthLargest{
				heap: MinHeap{
					array: []int{4, 5, 8},
				},
				k: 3,
			},
		},
	}

	Run_Tests(t, tests, func(test TestCase[input]) {
		actual := Constructor(test.Args.k, test.Args.nums)
		Equal(t, test.Expected, actual)
	})
}

func Test_KthLargestAdd(t *testing.T) {
	type input struct {
		k    int
		nums []int
	}

	tests := []TestCase[input]{
		{
			Name: "Add a number",
			Args: input{
				k:    3,
				nums: []int{4, 5, 8, 2},
			},
			Expected: 5,
		},
	}

	Run_Tests(t, tests, func(test TestCase[input]) {
		heap := Constructor(test.Args.k, test.Args.nums)
		actual := heap.Add(10)
		Equal(t, test.Expected, actual)
	})
}

func Test_KthLargestAddSeveral(t *testing.T) {
	type input struct {
		k    int
		nums []int
	}

	tests := []TestCase[input]{
		{
			Name: "Add 3 numbers",
			Args: input{
				k:    3,
				nums: []int{4, 5, 8, 2},
			},
			Expected: 8,
		},
	}

	Run_Tests(t, tests, func(test TestCase[input]) {
		heap := Constructor(test.Args.k, test.Args.nums)
		heap.Add(42)
		heap.Add(100)
		actual := heap.Add(6)
		Equal(t, test.Expected, actual)
	})
}
