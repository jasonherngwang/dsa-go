package main

import (
	"fmt"
	"math"
)

func main() {
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	var findMax func(node *TreeNode) int

	findMax = func(node *TreeNode) int {
		if node == nil {
			return math.MinInt
		}

		leftMaxVal := findMax(node.Left)
		rightMaxVal := findMax(node.Right)

		// Return the max out of itself and its two children
		if node.Val >= leftMaxVal && node.Val >= rightMaxVal {
			return node.Val
		}
		if leftMaxVal >= rightMaxVal {
			return leftMaxVal
		}
		return rightMaxVal
	}

	tree := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 8,
			},
			Right: &TreeNode{
				Val: 11,
			},
		},
	}

	fmt.Println(findMax(tree))

	tree2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 8,
			},
			Right: &TreeNode{
				Val: -11,
			},
		},
	}

	fmt.Println(findMax(tree2))
}
