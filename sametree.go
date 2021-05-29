type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func height(p *TreeNode) int {
	if p == nil {
		return 0
	}
	return 1 + height(p.Left) + height(p.Right)
}

func walk(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	var x = walk(p.Left, q.Left)
	var y = walk(p.Right, q.Right)
	return x && y
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if height(p) != height(q) {
		return false
	}
	return walk(p, q)
}