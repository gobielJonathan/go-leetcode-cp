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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var l = height(root.Left)
	var r = height(root.Right)
	return (r-l) <= 1 && (r-l) >= -1
}