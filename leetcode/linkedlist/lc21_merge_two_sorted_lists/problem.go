/*
Problem
Merge Two Sorted Lists
https://leetcode.com/problems/merge-two-sorted-lists/

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing
together the nodes of the first two lists.

Return the head of the merged linked list.

Inputs: The heads of 2 sorted linked lists
Outputs: The head of 1 new sorted linked list, formed from the two input lists

Rules, Requirements, Definitions
- Num of nodes: 0 to 50
  - Lists can be empty.
- Node values -100 to 100
  - Negative numbers and 0 are valid values.
- Lists are sorted ascending.
  - Output list should also be sorted ascending.
- An empty linked list is indicated by a null value for head.
- One empty? Return head of non-empty list.
- Both empty? Return head referencing null.

Questions
- Nodes in both lists with same value? Order doesn't matter.
- Will returned result list be completely separate from the input lists, i.e. no
  references to any nodes from the input lists?
  - No, the resulting list can reference nodes in either list.
  - We can also mutate the input lists in-place.

Examples, Test Cases, Edge Cases
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Input: list1 = [], list2 = []
Output: []

Input: list1 = [], list2 = [0]
Output: [0]

Data Structure
Linked list

Algorithm

Approach 1: Iterative
Time: O(N + M), where N and M are the input list lengths.
- Must traverse over every element in both lists.
Space: O(1) for a few pointers.

Steps
- Create empty listnode to serve as a "pre-head" for the result list. We will
  return the node AFTER the pre-head as the head of the result list.
  - We do this because for singly-linked lists we will have no way to get back
    to the start of the list after we finish building it.
- Create a pointer referencing the empty listnode. We will advance this after
  every step.
- We are given as inputs 2 pointers, pointing to the heads of the 2 input lists.
- Compare the values at the nodes, as long as neither is null.
  - If one lesser or equal:
    - Link the result pointer to this node.
    - Advance pointer for this list.
  - Else
    - Link the result pointer to the other list's node.
    - Advance pointer for the other list.
  - Advance the result pointer.
- Return the `next` property of the pre-head.

Approach 2: Recursion
Time: O(N + M) because we visit all nodes in both lists.
Space: O(N + M) for N + M stack frames.

Steps
- Base case: If either input list is empty, return the rest of the other one.
  - This means we have reached the end of one of the lists, so we tack on what's
    left of the non-empty list.
- Compare values at the nodes:
  - If first list's value is smaller, set that node's `next` to the return value
    of a recursive call, passing as argument the rest of the first list, and the
    entirety of the second list.
    - Return the node
  - Else, do the same using the second list.
*/

package leetcode

import "github.com/jasonherngwang/dsa-go/structures"

type ListNode = structures.ListNode

// Approach 1: Iterative
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	node := preHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}

	if list1 == nil {
		node.Next = list2
	} else {
		node.Next = list1
	}

	return preHead.Next
}

// Approach 2: Recursion
func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoListsRecursive(list1, list2.Next)
		return list2
	}
}
