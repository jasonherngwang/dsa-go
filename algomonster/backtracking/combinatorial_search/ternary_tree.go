/*
Find all root-to-leaf paths
A ternary tree has nodes with at most 3 children.

Return value: None; mutations only.
- Append the current node's value to the path traveled so far.
- Mutate the result array if we've reached the end of the path.

Pass state:
- The path traveled so far
- A reference to a global result variable, an array storing all the paths

Steps
- If leaf node (0-length children array):
  - Append val to path
  - Add path-so-far to result array

- Else
  - Recursive call on all children, passing path-so-far and result array ref.
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	type Node struct {
		val      int
		children []Node
	}

	ternaryTreePaths := func(root Node) []string {
		// Use pointer so reassignment is mutating for allPaths
		var dfs func(node Node, pathSoFar []string, allPaths *[]string)
		dfs = func(node Node, pathSoFar []string, allPaths *[]string) {
			// Base case: leaf node (no children), so we update allPaths
			if len(node.children) == 0 {
				pathSoFar = append(pathSoFar, strconv.Itoa(node.val))
				*allPaths = append(*allPaths, strings.Join(pathSoFar, "->"))
				return
			}

			for _, child := range node.children {
				// If not leaf node, add self and continue recursing
				if len(node.children) > 0 {
					pathSoFar = append(pathSoFar, strconv.Itoa(node.val)) // Push
					dfs(child, pathSoFar, allPaths)
					// Pop, so we start fresh with the next child
					pathSoFar = pathSoFar[:len(pathSoFar)-1]
				}
			}
		}

		allPaths := &[]string{}
		initialPath := []string{}
		// Guard clause: root is null
		if len(root.children) != 0 {
			dfs(root, initialPath, allPaths)
		}

		return *allPaths
	}

	tree := Node{
		val: 1,
		children: []Node{
			Node{
				val: 2,
				children: []Node{
					Node{
						val: 3,
					},
					// Node{
					// 	val: 42,
					// },
				},
			},
			Node{
				val:      4,
				children: []Node{},
			},
			Node{
				val:      6,
				children: []Node{},
			},
		},
	}

	fmt.Println(ternaryTreePaths(tree))
}
