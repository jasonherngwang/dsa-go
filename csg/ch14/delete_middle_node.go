/*
Given a reference to a node in the middle of a singly-linked list, delete it.

1 2 3 4 5 -> nil
    |

Copy the next node's value into the current node, overwriting it. This
effectively "shifts" all nodes over by 1. Then we need to remove the last node
by setting the second-to-last node's next to nil.

1 2 4 4 5 -> nil
    |

1 2 4 5 5 -> nil
      |

1 2 4 5 -> nil
      |
*/

package main

import "fmt"

func main() {
	type Node struct {
		val  int
		next *Node
	}

	deleteNode := func(node *Node) {
		for node.next.next != nil {
			node.val = node.next.val
			node = node.next
		}
		node.val = node.next.val
		node.next = nil
	}

	middleNode := &Node{
		val: 3,
		next: &Node{
			val: 4,
			next: &Node{
				val:  5,
				next: nil,
			},
		},
	}

	list := &Node{
		val: 1,
		next: &Node{
			val:  2,
			next: middleNode,
		},
	}

	curr := list
	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.next
	}

	deleteNode(middleNode)

	curr = list
	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.next
	}

}
