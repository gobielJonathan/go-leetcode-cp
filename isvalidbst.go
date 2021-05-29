
func isLeaf(r *TreeNode) bool {
	if r == nil {
		return false
	}
	return r.Left == nil && r.Right == nil
}
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var temp = root.Right
	if temp != nil {
		for temp != nil && !isLeaf(temp) {
			temp = temp.Left
		}
		if temp == nil {
			return false
		}
		return temp.Val > root.Val
	}
	temp = root.Left
	if temp != nil {
		for temp != nil && !isLeaf(temp) {
			temp = temp.Right
		}
		if temp == nil {
			return false
		}
		return temp.Val < root.Val
	}
	return isValidBST(root.Left) && isValidBST(root.Right)
}