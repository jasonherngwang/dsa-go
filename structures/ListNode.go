package structures

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
}

// Utilities

// Print data of all nodes, space-separated (given LinkedList)
func (l LinkedList) PrintListData() {
	curr := l.Head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Printf("\n")
}

// Print data of all nodes, space-separated (given ListNode)
func (head *ListNode) PrintData() {
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Printf("\n")
}

// Return list length
func (l LinkedList) GetLength() int {
	length := 0
	curr := l.Head

	for curr != nil {
		length++
		curr = curr.Next
	}

	return length
}

// Create sample list, for use in testing
func CreateLinkedList() LinkedList {
	list := LinkedList{}
	for n := 1; n <= 5; n++ {
		node := &ListNode{Val: n}
		list.AppendNode(node)
	}

	return list
}

// Addition

// Prepend node to list O(1)
func (l *LinkedList) PrependNode(newHead *ListNode) {
	prevHead := l.Head
	newHead.Next = prevHead
	l.Head = newHead
}

// Append node to list O(N)
func (l *LinkedList) AppendNode(n *ListNode) {
	last := l.Head

	// Empty list
	if last == nil {
		l.Head = n
		return
	}

	// Find last node
	for last.Next != nil {
		last = last.Next
	}
	last.Next = n
}

// Insertion

// Insert value (as node) at specified index O(N)
func (l *LinkedList) InsertValueAtIndex(value int, index int) {
	newNode := &ListNode{Val: value}

	// Empty list
	if l.Head == nil {
		l.Head = newNode
		return
	}

	// Edge case: Prepend at beginning of list
	if index == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		return
	}

	curr := l.Head // newNode will be inserted before this
	currIndex := 0
	// Traverse to specified index
	for currIndex < (index - 1) {
		curr = curr.Next
		currIndex++

		// Provided index is out-of-bounds
		if curr == nil {
			return
		}
	}

	// Insert new node at index
	newNode.Next = curr.Next
	curr.Next = newNode
}

// Retrieval

// Return pointer to node at specified index O(N)
func (l LinkedList) GetNodeAtIndex(index int) (*ListNode, bool) {
	curr := l.Head
	currIndex := 0

	for currIndex < index {
		curr = curr.Next
		currIndex++

		// Provided index is out-of-bounds
		if curr == nil {
			return nil, true
		}
	}

	return curr, false
}

// Search

// Return pointer to first node having specified value O(N)
func (l LinkedList) GetNodeWithValue(value int) (*ListNode, bool) {
	curr := l.Head

	for curr != nil {
		if curr.Val == value {
			return curr, false
		}
		curr = curr.Next
	}

	// No match
	return nil, true
}

// Delete first instance of value O(N)
func (l *LinkedList) DeleteFirstMatch(value int) {
	// Empty list
	if l.Head == nil {
		return
	}

	// First node matches; delete it
	if l.Head.Val == value {
		l.Head = l.Head.Next
		return
	}

	prev := l.Head

	// Traverse until first match, then skip it
	for prev.Next.Val != value {
		// Reached end of list
		if prev.Next.Next == nil {
			return
		}
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
}

// Delete all instances of value O(N)
func (l *LinkedList) DeleteAllMatches(value int) {
	// Empty list
	if l.Head == nil {
		return
	}

	// One node
	if l.Head.Next == nil {
		if l.Head.Val == value {
			l.Head = nil
		}
		return
	}

	prev := l.Head
	// Traverse until the last node, skipping any matches
	for prev.Next != nil {
		if prev.Next.Val == value {
			prev.Next = prev.Next.Next
		}
		prev = prev.Next
	}
}

// Deletion

// Delete node at specified index
func (l *LinkedList) DeleteNodeAtindex(index int) {
	if index == 0 {
		l.Head = l.Head.Next
		return
	}

	curr := l.Head
	currIndex := 0

	for currIndex < index-1 {
		curr = curr.Next
		currIndex++

		// Provided index is out-of-bounds
		if curr.Next == nil {
			return
		}
	}

	curr.Next = curr.Next.Next
}
