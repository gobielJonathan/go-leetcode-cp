
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

const MAX_U = ^uint(0)
const MIN_U = 0
const MAX_I = int(MAX_U >> 1)
const MIN_I = -MAX_I - 1

func findLargest(root *TreeNode, target int, level int) int {
	if root == nil {
		return MIN_I
	}
	var l = findLargest(root.Left, target, level+1)
	if target == level {
		return root.Val
	}
	var r = findLargest(root.Right, target, level+1)
	if r > l {
		return r
	}
	return l
}
func largestValues(root *TreeNode) []int {
	var h = height(root)
	var temp = []int{}
	for i := 1; i <= h; i++ {
		temp = append(temp, findLargest(root, i, 1))
	}
	return temp
}