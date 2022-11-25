/*
Problem
Remove Nth Node From End of List
https://leetcode.com/problems/remove-nth-node-from-end-of-list/

Given the head of a linked list, remove the nth node from the end of the list
and return its head.

Inputs: Head of a linked list
Outputs: Head of the same list, with Nth node from end removed

Rules, Requirements, Definitions
- Linked list will have 1-30 nodes.
- N will be between 1 and the number of nodes (no out-of-bounds values).

Examples, Test Cases, Edge Cases
Generic case: N is in the middle
- Input: head = [1,2,3,4,5], n = 2
- Output: [1,2,3,5]

Case: List of 1 node; N = 1.
- Input: head = [1], n = 1
- Output: []
  - Head points to null.

Case: List of M nodes; N = M.
- Head points to what was previously the 2nd node.

Case: List of M nodes; N = 1.
- Input: head = [1,2], n = 1
- Output: [1]
  - Last node points to null.


Data Structure
Linked list

Algorithm

Time: O(N) to traverse all N nodes.
Space: O(1) for a few pointers.

Step 1: Find Nth from last node.
Use two pointers, with N nodes between Left and Right
- Create a dummy head linked to head. This will be our return value.
- Initialize Left to dummy.
- Initialize Right to dummy and advance it N+1 times, so there is a gap of size
  N between the two pointers.
- While Right is not null, advance both by 1 node.
- Now Left is the (N+1)th node from the end.

Step 2: Remove the node.
- Set left.next to left.next.next, to remove the node.
  - Handles edge case where left = dummy. It skips the first node. This is OK
	because our return value is dummy.next, not dummy itself.

Step 3: Return dummy.next
*/

package leetcode

import "github.com/jasonherngwang/dsa-go/structures"

type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{Next: head}
	left := dummy
	right := dummy

	// Create gap of size n between left and right
	for step := 0; step <= n; step++ {
		right = right.Next
	}

	// Advance right to end
	for right != nil {
		left = left.Next
		right = right.Next
	}

	// Remove node
	left.Next = left.Next.Next
	return dummy.Next
}
