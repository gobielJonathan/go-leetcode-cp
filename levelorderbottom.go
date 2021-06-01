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
func takeWithLevel(r *TreeNode, target int, level int) []int {
	if r == nil {
		return nil
	}
	var x = takeWithLevel(r.Left, target, level+1)
	if target == level {
		return []int{r.Val}
	}
	var y = takeWithLevel(r.Right, target, level+1)
	return append(x, y...)
}
func levelOrderBottom(root *TreeNode) [][]int {
	var h = height(root)
	var temp = [][]int{}
	for i := h; i >= 1; i-- {
		temp = append(temp, takeWithLevel(root, i, 1))
	}
	return temp
}