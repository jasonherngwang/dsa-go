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

	// Uses global variable which is reassigned as necessary
	max := math.MinInt

	var findMax func(node *TreeNode)

	findMax = func(node *TreeNode) {
		if node == nil {
			return
		}

		findMax(node.Left)
		findMax(node.Right)

		if node.Val > max {
			max = node.Val
		}
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

	findMax(tree)
	fmt.Println(max)

	// Reset
	max = math.MinInt

	tree = &TreeNode{
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

	findMax(tree)
	fmt.Println(max)
}
