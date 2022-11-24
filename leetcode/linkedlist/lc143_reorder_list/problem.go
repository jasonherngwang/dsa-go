/*
Problem
Reorder Linked List
https://leetcode.com/problems/reorder-list/

You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln-1 → Ln

Reorder the list to be on the following form:

L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …

You may not modify the values in the list's nodes.
Only nodes themselves may be changed.

Inputs: Head of a singly-linked list
Outputs: Nothing. Re-order list in-place.

Rules, Requirements, Definitions
- Alternate between start and end, not including any that have been visited.
- Remember to point the last node's `next` to null.

Questions


Examples, Test Cases, Edge Cases


Data Structure
Linked list

Algorithm

Approach 1:
- Find middle.
- Split into halves.
- Reverse second half.
- Merge halves, alternating.

Time: O(N)
- Find middle node: O(N)
- Reverse second half: O(N/2)
- Merge lists: O(N/2)
Space: O(1). Re-order list in-place. Use some pointers.

Step 1: Split into halves AKA Find the middle of the list
Use Slow and Fast pointers:
- "Slow" means it advances 1 node at at time. "Fast" advances 2 at at time.
- When Fast reaches the end, Slow should be at the middle.
- Slow pointer starts at first node (head).
- Fast pointer starts at second node (head.next).
- While fast not null (odd num of nodes) and fast.next not null (even num):
  - Advance slow by 1 (slow.next).
  - Advance fast by 2 (fast.next.next).
- `slow` now references the middle (odd num), or left-of-middle (even num).
  - For odd num of nodes, the first half will be longer by 1.
- Create pointer `second` referencing slow.next, the head of the second half.

Step 2: Reverse second half
- Three pointers: prev = null, curr = head, next
- While curr is not null:
  - Set next to curr.next.
  - Set curr to prev.
  - Set prev to curr.
  - Set curr to next.
- Return prev.

Step 3: Merge linked lists, alternating
- We have two pointers: head of first half, head of the reversed second half.
- Merge halves. Second half is shorter, so while second pointer is not null:
  - Link first half's node to second half's node.
  - Link second half's node to first half's next node.
  - Advance both pointers.

Approach 2: Array
- Convert linked list to array of values.
- Find array length, and halfway index (right-of-middle for even num).
- Create two pointers starting at the list head: curr and last.
- Iterate to halfway (left of middle). At each index:
	- Calculate the "mirror" index.
	- Create a new node containing that mirror node's value.
		- Set the new node's Next to curr.Next
	- Set curr.Next to this new node
	- Advance curr to the new node's Next (curr's original Next)
	- Reassign last to this new node
- If there are an even number of nodes, the new node will be last. If even, the
  curr node will be the last. Set the last node's Next to nil.
*/

package leetcode

import "github.com/jasonherngwang/dsa-go/structures"

type ListNode = structures.ListNode

// Approach 1
func reorderList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow := head
	fast := head

	// Find middle
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Detach second half
	curr := slow.Next
	slow.Next = nil

	// Reverse latter half
	var prev *ListNode // nil pointer
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// Merge lists
	first := head
	second := prev

	// For odd num of nodes, second half is shorter
	for second != nil {
		firstNext := first.Next
		secondNext := second.Next

		first.Next = second
		second.Next = firstNext

		first = firstNext
		second = secondNext
	}

	return head
}

// Approach 2
func listToArray(head *ListNode) []int {
	array := []int{}
	curr := head

	for curr != nil {
		array = append(array, curr.Val)
		curr = curr.Next
	}

	return array
}

func reorderListUsingArray(head *ListNode) *ListNode {
	curr := head
	new := head

	array := listToArray(head)
	arrayLen := len(array)
	midIndex := arrayLen / 2

	for i := 0; i < midIndex; i++ {
		mirrorIndex := arrayLen - 1 - i
		mirrorVal := array[mirrorIndex]
		newNode := &ListNode{Val: mirrorVal, Next: curr.Next}

		curr.Next = newNode
		curr = newNode.Next
		new = newNode
	}

	if arrayLen%2 == 0 {
		new.Next = nil
	} else {
		curr.Next = nil
	}

	return head
}
