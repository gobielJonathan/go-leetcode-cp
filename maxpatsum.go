func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var x = maxPathSum(root.Left)
	var y = maxPathSum(root.Right)
	if (root.Val + x + y) < x {
		return x
	}
	if (root.Val + x + y) < y {
		return y
	}
	return root.Val + x + y
}