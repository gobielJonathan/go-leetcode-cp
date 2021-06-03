package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isleaf(r *TreeNode) bool {
	return r.Left == nil && r.Right == nil
}

func getpaths(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var l = getpaths(root.Left)
	var r = getpaths(root.Right)

	return l + r + " "
}
func sumNumbers(root *TreeNode) int {
	fmt.Println(getpaths(root))
	return 0
}

func main() {
	// https://leetcode.com/problems/sum-root-to-leaf-numbers/
	// https://leetcode.com/problems/add-one-row-to-tree/
	var r = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	fmt.Println(sumNumbers(r))
}
