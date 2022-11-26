package structures

type DLLNode struct {
	Val  int
	Next *DLLNode
	Prev *DLLNode
}

type DoublyLinkedLinkedList struct {
	Head *DLLNode
	Tail *DLLNode
}
