package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTally(t *testing.T) {
	t.Run("tally slice of ints", func(t *testing.T) {
		ints := []int{1, 2, 2, 3, 3, 3}

		actual := Tally[int](ints)
		expected := map[int]int{
			1: 1,
			2: 2,
			3: 3,
		}
		assert.Equal(t, expected, actual)
	})
}
