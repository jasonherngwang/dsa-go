package utils

import (
	"testing"

	"github.com/ethanjweiner/go-test-helper/testhelper"
	"github.com/stretchr/testify/assert"
)

func Test_insertionSort(t *testing.T) {
	type input struct {
		nums []int
	}

	tests := []testhelper.TestCase[input]{
		{
			Name: "Insertion sort",
			Args: input{
				nums: []int{3, -1, 2, 0, 1},
			},
			Expected: []int{-1, 0, 1, 2, 3},
		},
		{
			Name: "Insertion sort 1 elem",
			Args: input{
				nums: []int{42},
			},
			Expected: []int{42},
		},
		{
			Name: "Insertion sort empty slice",
			Args: input{
				nums: []int{},
			},
			Expected: []int{},
		},
	}

	testhelper.Run_Tests(t, tests, func(test testhelper.TestCase[input]) {
		insertionSort(test.Args.nums)
		assert.Equal(t, test.Expected, test.Args.nums)
	})
}

func Test_selectionSort(t *testing.T) {
	type input struct {
		nums []int
	}

	tests := []testhelper.TestCase[input]{
		{
			Name: "Selection sort",
			Args: input{
				nums: []int{3, -1, 2, 0, 1},
			},
			Expected: []int{-1, 0, 1, 2, 3},
		},
		{
			Name: "Selection sort 1 elem",
			Args: input{
				nums: []int{42},
			},
			Expected: []int{42},
		},
		{
			Name: "Selection sort empty slice",
			Args: input{
				nums: []int{},
			},
			Expected: []int{},
		},
	}

	testhelper.Run_Tests(t, tests, func(test testhelper.TestCase[input]) {
		selectionSort(test.Args.nums)
		assert.Equal(t, test.Expected, test.Args.nums)
	})
}

func Test_bubbleSort(t *testing.T) {
	type input struct {
		nums []int
	}

	tests := []testhelper.TestCase[input]{
		{
			Name: "Bubble sort",
			Args: input{
				nums: []int{3, -1, 2, 0, 1},
			},
			Expected: []int{-1, 0, 1, 2, 3},
		},
		{
			Name: "Bubble sort",
			Args: input{
				nums: []int{3, 0, -1, 1, 2},
			},
			Expected: []int{-1, 0, 1, 2, 3},
		},
		{
			Name: "Bubble sort 1 elem",
			Args: input{
				nums: []int{42},
			},
			Expected: []int{42},
		},
		{
			Name: "Bubble sort empty slice",
			Args: input{
				nums: []int{},
			},
			Expected: []int{},
		},
	}

	testhelper.Run_Tests(t, tests, func(test testhelper.TestCase[input]) {
		bubbleSort(test.Args.nums)
		assert.Equal(t, test.Expected, test.Args.nums)
	})
}

func Test_mergeSort(t *testing.T) {
	type input struct {
		arr []int
	}

	tests := []testhelper.TestCase[input]{
		{
			Name: "Merge sort",
			Args: input{
				arr: []int{3, -1, 2, 0, 1},
			},
			Expected: []int{-1, 0, 1, 2, 3},
		},
		{
			Name: "Merge sort with duplicates",
			Args: input{
				arr: []int{3, 0, -1, -1, 0},
			},
			Expected: []int{-1, -1, 0, 0, 3},
		},
		{
			Name: "Merge sort 1 elem",
			Args: input{
				arr: []int{42},
			},
			Expected: []int{42},
		},
		{
			Name: "Merge sort empty slice",
			Args: input{
				arr: []int{},
			},
			Expected: []int{},
		},
	}

	testhelper.Run_Tests(t, tests, func(test testhelper.TestCase[input]) {
		actual := mergeSort[int](test.Args.arr)
		assert.Equal(t, test.Expected, actual)
	})
}

func Test_quickSort(t *testing.T) {
	type input struct {
		arr []int
	}

	tests := []testhelper.TestCase[input]{
		{
			Name: "Quick sort",
			Args: input{
				arr: []int{3, -1, 2, 0, 1},
			},
			Expected: []int{-1, 0, 1, 2, 3},
		},
		{
			Name: "Quick sort with duplicates",
			Args: input{
				arr: []int{3, 0, -1, -1, 0},
			},
			Expected: []int{-1, -1, 0, 0, 3},
		},
		{
			Name: "Quick sort 1 elem",
			Args: input{
				arr: []int{42},
			},
			Expected: []int{42},
		},
		{
			Name: "Quick sort empty slice",
			Args: input{
				arr: []int{},
			},
			Expected: []int{},
		},
	}

	testhelper.Run_Tests(t, tests, func(test testhelper.TestCase[input]) {
		quickSort[int](test.Args.arr)
		assert.Equal(t, test.Expected, test.Args.arr)
	})
}
