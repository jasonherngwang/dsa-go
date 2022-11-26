package structures

import (
	"testing"
)

func TestInsert(t *testing.T) {
	t.Run("Insert into hash table", func(t *testing.T) {
		hashTable := Init()
		list := []string{
			"Fido",
			"Spot",
			"Rover",
		}

		for _, v := range list {
			hashTable.Insert(v)
		}

		actual := hashTable.Exists("Spot")
		expected := true

		if actual != expected {
			t.Errorf("Expected: %t, received %t", expected, actual)
		}
	})
}
