package structures

import "fmt"

// MaxHeap
type MaxHeap struct {
	array []int
}

// Add element
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Remove largest element
func (h *MaxHeap) Extract() int {
	// Remove root
	max := h.array[0]
	lastIndex := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("Array length is 0; cannot extract")
		return -1
	}

	// Move last node to root
	h.array[0] = h.array[lastIndex]
	// Remove last node
	h.array = h.array[:lastIndex]

	// Push new root down
	h.maxHeapifyDown(0)

	return max
}

// Heapify down
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		// In complete tree, nodes are pushed to left
		if l == lastIndex || h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		// swap
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			// Already in correct place
			return
		}
	}
}

// Heapify up
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index) // Swap larger node with smaller node
		index = parent(index)        // Update pointer
	}
}

// Index of parent node
func parent(index int) int {
	return (index - 1) / 2
}

// Index of left child
func left(index int) int {
	return index*2 + 1
}

// Index of right child
func right(index int) int {
	return index*2 + 2
}

// Swap nodes
func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}
