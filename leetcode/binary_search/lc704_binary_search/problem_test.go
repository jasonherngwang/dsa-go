package leetcode

import (
	"testing"

	"github.com/ethanjweiner/go-test-helper/testhelper"
	"github.com/stretchr/testify/assert" // Or your library of choice
)

func Test_binarySearch(t *testing.T) {
	type input struct {
		nums   []int
		target int
	}

	tests := []testhelper.TestCase[input]{
		{
			Name: "Find number that is present",
			Args: input{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			Expected: 4,
		},
		{
			Name: "Find number that is not present",
			Args: input{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			Expected: -1,
		},
	}

	testhelper.Run_Tests(t, tests, func(test testhelper.TestCase[input]) {
		actual := binarySearch(test.Args.nums, test.Args.target)
		assert.Equal(t, test.Expected, actual)
	})
}
