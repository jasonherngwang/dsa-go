/*
Problem
------------------------------------------
K Closest Points to Origin
https://leetcode.com/problems/k-closest-points-to-origin/

Given an array of points where points[i] = [xi, yi] represents a point on the
X-Y plane and an integer k, return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance
(i.e., √(x1 - x2)2 + (y1 - y2)2).

You may return the answer in any order. The answer is guaranteed to be unique
(except for the order that it is in).

Constraints:

1 <= k <= points.length <= 104
-104 < xi, yi < 104

Inputs:
- Array of 2-element arrays (points)
- Integer k
Outputs: Array of length k, of 2-element arrays (points)

Rules, Requirements, Definitions
------------------------------------------
- Can there be multiple points of the same distance from (0,0)? No.


Examples, Test Cases, Edge Cases
------------------------------------------
Case:
- Input: [[1,3],[-2,2]], k = 1
- Output: [[-2,2]]

Case:
- Input: [[3,3],[5,-1],[-2,4]], k = 2
- Output: [[3,3],[-2,4]]

Data Structure
------------------------------------------
Heap

Algorithm
------------------------------------------

Time:
Space:

Steps
- Initialize min heap where each node stores 2 pieces of data:
  - (x, y) coord
  - Computed distance from (0, 0)

Helper function: Absolute distance from (0, 0)
- Return square root of sum of squares of x and y
*/

package leetcode

import (
	"fmt"
	"math"
)

type MinHeap struct {
	array []PointData
}

type PointData struct {
	coords   []int
	distance float64
}

func (h *MinHeap) Insert(point PointData) {
	h.array = append(h.array, point)
	h.minHeapifyUp(len(h.array) - 1)
}
func (h *MinHeap) Extract() PointData {
	min := h.array[0]
	lastIndex := len(h.array) - 1

	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	h.minHeapifyDown(0)
	return min
}

func (h *MinHeap) minHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex || h.array[l].distance < h.array[r].distance {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index].distance > h.array[childToCompare].distance {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func (h *MinHeap) minHeapifyUp(index int) {
	for h.array[parent(index)].distance > h.array[index].distance {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}
func left(index int) int {
	return index*2 + 1
}
func right(index int) int {
	return index*2 + 2
}
func (h *MinHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func distanceFromOrigin(coords []int) float64 {
	x, y := float64(coords[0]), float64(coords[1])
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}

func kClosest(points [][]int, k int) [][]int {
	m := &MinHeap{}
	result := [][]int{}

	for _, point := range points {
		distance := distanceFromOrigin(point)
		m.Insert(PointData{
			coords:   point,
			distance: distance,
		})
	}

	fmt.Println(m)

	for i := 0; i < k; i++ {
		result = append(result, m.Extract().coords)
	}

	return result
}
