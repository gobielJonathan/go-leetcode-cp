package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func height(p *TreeNode) int {
	if p == nil {
		return 0
	}
	var l = height(p.Left)
	var r = height(p.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
func getValBaseLevel(root *TreeNode, target int, level int) []int {
	var temp = []int{}
	var x = getValBaseLevel(root.Left, target, level+1)
	if target == level {
		return append(temp, root.Val)
	}
	var y = getValBaseLevel(root.Right, target, level+1)
	return append(x, y...)
}
func widthOfBinaryTree(root *TreeNode) int {
	var h = height(root)
	var count = 0
	for i := 1; i <= h; i++ {
		var l = getValBaseLevel(root, i, 1)
		if len(l) > count {
			count = len(l)
		}
	}
	return count
}

func main() {
	// https://leetcode.com/problems/maximum-width-of-binary-tree/
}
