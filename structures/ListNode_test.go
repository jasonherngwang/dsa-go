package structures

import (
	// "fmt"
	"testing"
)

func TestGetLength(t *testing.T) {
	t.Run("Returns correct length", func(t *testing.T) {
		list := CreateLinkedList()

		actual := list.GetLength()
		expected := 5

		if actual != expected {
			t.Errorf("Expected: %d, received %d", expected, actual)
		}
	})
}

func TestInsertValueAtIndex(t *testing.T) {
	t.Run("Successfully inserts value", func(t *testing.T) {
		list := CreateLinkedList()

		list.InsertValueAtIndex(42, 2)
		actual := list.Head.Next.Next.Val
		expected := 42

		if actual != expected {
			t.Errorf("Expected: %d, received %d", expected, actual)
		}
	})
}

func TestGetNodeAtIndex(t *testing.T) {
	t.Run("Finds valid value", func(t *testing.T) {
		list := CreateLinkedList()

		actual, err := list.GetNodeAtIndex(0)
		if err {
			t.Errorf("Value not found")
			return
		}
		expected := 1

		if actual.Val != expected {
			t.Errorf("Expected: %d, received %d", expected, actual.Val)
		}
	})
}

func TestGetNodeWithValue(t *testing.T) {
	t.Run("Finds node containing valid value", func(t *testing.T) {
		list := CreateLinkedList()

		actual, err := list.GetNodeWithValue(3)
		if err {
			t.Errorf("Value not found")
			return
		}
		expected := list.Head.Next.Next

		if actual != expected {
			t.Errorf("Expected: %d, received %d", expected.Val, actual.Val)
		}
	})
}

func TestDeleteFirstMatch(t *testing.T) {
	t.Run("Deletes first instance of valid value", func(t *testing.T) {
		list := CreateLinkedList()

		list.DeleteFirstMatch(2)

		actual := list.GetLength()
		expected := 4

		if actual != expected {
			t.Errorf("Expected: %d, received %d", expected, actual)
		}
	})
}

func TestDeleteAllMatches(t *testing.T) {
	t.Run("Deletes all instances of valid value", func(t *testing.T) {
		list := CreateLinkedList()
		// Add duplicate values
		list.AppendNode(&ListNode{Val: 2})
		list.AppendNode(&ListNode{Val: 3})

		list.DeleteAllMatches(2)

		actual := list.GetLength()
		expected := 5

		if actual != expected {
			t.Errorf("Expected: %d, received %d", expected, actual)
		}
	})
	t.Run("Doesn't delete if value not in list", func(t *testing.T) {
		list := CreateLinkedList()

		list.DeleteAllMatches(42)

		actual := list.GetLength()
		expected := 5

		if actual != expected {
			t.Errorf("Expected: %d, received %d", expected, actual)
		}
	})
	t.Run("Doesn't delete if list is already empty", func(t *testing.T) {
		list := LinkedList{}

		list.DeleteAllMatches(1)
		list.PrintListData()

		actual := list.GetLength()
		expected := 0

		if actual != expected {
			t.Errorf("Expected: %d, received %d", expected, actual)
		}
	})
}
