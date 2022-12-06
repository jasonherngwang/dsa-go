/*
Serialize and Deserialize Binary Tree
- Store serialized data as a string (global var)

Info needed at each node to make a decision: Just the current value.
Pass state down: None needed

Serialize
DFS pre-order (curr, left, right), gathering values into a comma-separated
string.
- Concatenate current value, or "X" if nil

       1
  2         5
3   4     6   X

[1, 2, 3, X, X, 4, X, X, 5, 6, X, X, X]

             1
      2           6
  3       X     X   X
X   X
[1, 2, 3, X, X, X, 6, X, X]

Deserialize
- Split string by comma into array of tokens
- DFS pre-order
  - Retrieve latest token and increment global index var tracking our position
    in the token array.
  - If "X" return nil; this will be the child of the parent node.
  - Else create a new Node.
  - Check if end of token array reached. If so, the final step is returning the
    new node.
  - Set left and right to recursive calls of those children
  - Return new node to be assigned as child of the parent node.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	type Node struct {
		val   int
		left  *Node
		right *Node
	}

	serialize := func(root *Node) string {
		result := ""

		var dfs func(root *Node)
		dfs = func(root *Node) {
			if root == nil {
				result += ",X"
				return
			}

			if len(result) > 0 {
				result += ","
			}
			result += strconv.Itoa(root.val)

			dfs(root.left)
			dfs(root.right)

		}

		dfs(root)

		return result
	}

	deserialize := func(root string) *Node {
		// Track current location in the token array
		i := 0

		var dfs func(tokens []string) *Node

		dfs = func(tokens []string) *Node {
			token := tokens[i]
			i++

			if token == "X" {
				return nil
			}

			val, err := strconv.Atoi(token)
			if err != nil {
				return nil
			}

			node := &Node{
				val: val,
			}

			// Completed processing all tokens
			if i == len(tokens) {
				return node
			}

			// Recurse until both children are assigned to nil (leaf node)
			node.left = dfs(tokens)
			node.right = dfs(tokens)

			return node
		}

		return dfs(strings.Split(root, ","))
	}

	tree := &Node{
		val: 1,
		left: &Node{
			val: 2,
			left: &Node{
				val: 3,
			},
		},
		right: &Node{
			val: 6,
		},
	}
	s := serialize(tree)
	fmt.Println(s)
	d := deserialize(s)
	fmt.Println(d)

	var printTreePreorder func(root *Node)
	printTreePreorder = func(root *Node) {
		if root == nil {
			fmt.Println("X")
			return
		}
		fmt.Println(root.val)
		printTreePreorder(root.left)
		printTreePreorder(root.right)
	}
	printTreePreorder(d)
}
