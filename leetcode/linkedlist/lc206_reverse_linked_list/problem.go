/*
Reverse Linked List
https://leetcode.com/problems/reverse-linked-list/

Problem
------------------------------------------
Given the head of a singly linked list, reverse the list, and return the
reversed list.

Inputs: The head, a pointer to the first node of a linked list
Outputs: A pointer to the first node of the reversed linked list

Rules, Requirements, Definitions
- Reverse the list in-place

Edge Cases
- Empty list (head points to nil)? Return pointer to nil.


Examples, Test Cases
------------------------------------------
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Input: head = []
Output: []

Data Structure
------------------------------------------
Linked list

Algorithm
------------------------------------------

Approach 1: Iterative
Pointers to current, prev, next. Point current to prev, and advance all 3 to the
next node.

Time: O(N) for 1 traversal
Space: O(1). Only using a few variables.
Steps:
- Initialize `curr` to head.
- Initialize `prev` to nil.

- While curr is not nil (traveled past end of list):
  - Initialize temporary variable `next` to curr's next.
  - Reassign curr's next to prev.
  - Reassign prev to current.
  - Reassign current to `next`.

- Return prev (the new head).
*/

package leetcode

import "github.com/jasonherngwang/dsa-go/structures"

type ListNode = structures.ListNode

// Approach 1: Iterative
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
