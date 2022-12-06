package main

import "fmt"

func main() {
	type TreeNode struct {
		Val   string
		Left  *TreeNode
		Right *TreeNode
	}

	indentPerLevel := "  "

	var dfsPrettyPrint func(node *TreeNode, indentLevel string)
	dfsPrettyPrint = func(node *TreeNode, indentLevel string) {
		if node == nil {
			return
		}

		currentIndentLevel := indentLevel + indentPerLevel
		fmt.Println(currentIndentLevel + node.Val)
		dfsPrettyPrint(node.Left, currentIndentLevel)
		dfsPrettyPrint(node.Right, currentIndentLevel)
	}

	dirStructure := &TreeNode{
		Val: "/",
		Left: &TreeNode{
			Val:  "Foo",
			Left: &TreeNode{Val: "Baz"},
		},
		Right: &TreeNode{
			Val: "Bar",
		},
	}

	dfsPrettyPrint(dirStructure, "")
}
