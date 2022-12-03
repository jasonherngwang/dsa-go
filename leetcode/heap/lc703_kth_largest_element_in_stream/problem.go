/*
Problem
------------------------------------------
Kth Largest Element in a Stream
https://leetcode.com/problems/kth-largest-element-in-a-stream/

Design a class to find the kth largest element in a stream. Note that it is the
kth largest element in the sorted order, not the kth distinct element.

Implement KthLargest class:
KthLargest(int k, int[] nums) Initializes the object with the integer k and the
stream of integers nums.
int add(int val) Appends the integer val to the stream and returns the element
representing the kth largest element in the stream.

Inputs:
- k (int)
- Array of initial numbers
Input for each Add operation:
- 1 integer to be added
Outputs of each Add operation:
- 1 integer, the Kth largest out of all the data gathered so far.

Rules, Requirements, Definitions
------------------------------------------
- Duplicates are allowed


Examples, Test Cases, Edge Cases
------------------------------------------
Input
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output
[null, 4, 5, 5, 8, 8]

Explanation
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8

Data Structure
------------------------------------------
Max heap

Algorithm
------------------------------------------
Use a min heap maintained at size k.

Time:
- Create heap: O(N)
- Add number: O(log N) for heapify up
- Extract number: O(log N) for heapify down
Space: O(k) for heap

Steps
- Initialize a min heap and insert initial values, heapifying up each time.
- Truncate the array to size k, by popping elements and heapifying down.
- Add operation:
  - Insert value and heapify up.
  - If length > k, pop smallest element.
  - Read and return smallest element (root).
*/

package leetcode

// MinHeap
type MinHeap struct {
	array []int
}

// Add element
func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.minHeapifyUp(len(h.array) - 1)
}

// Remove smallest element
func (h *MinHeap) Extract() int {
	// Remove root
	min := h.array[0]
	lastIndex := len(h.array) - 1

	// Move last node to root
	h.array[0] = h.array[lastIndex]
	// Remove last node
	h.array = h.array[:lastIndex]

	// Push new root down
	h.minHeapifyDown(0)

	return min
}

// Heapify down
func (h *MinHeap) minHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		// In complete tree, nodes are pushed to left
		if l == lastIndex || h.array[l] < h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		// swap
		if h.array[index] > h.array[childToCompare] {
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
func (h *MinHeap) minHeapifyUp(index int) {
	for h.array[parent(index)] > h.array[index] {
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
func (h *MinHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

type KthLargest struct {
	heap MinHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	minHeap := KthLargest{
		heap: MinHeap{
			array: []int{},
		},
		k: k,
	}

	for _, n := range nums {
		minHeap.heap.Insert(n)
	}

	for len(minHeap.heap.array) > k {
		minHeap.heap.Extract()
	}

	return minHeap
}

func (this *KthLargest) Add(val int) int {
	this.heap.Insert(val)
	if len(this.heap.array) > this.k {
		this.heap.Extract()
	}
	return this.heap.array[0]
}
