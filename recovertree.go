type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var l = getMin(root.Left)
	var r = getMin(root.Right)

	if root.Left == nil && root.Right == nil {
		return root
	}

	var min = root
	if root.Left == nil {
		if min.Val > root.Right.Val {
			min = root.Right
		}
	}
	if root.Right == nil {
		if min.Val > root.Left.Val {
			min = root.Left
		}
	}
	if l != nil && l.Val < min.Val {
		return l
	}
	if r != nil && r.Val < min.Val {
		return r
	}
	return min
}

func getMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var l = getMax(root.Left)
	var r = getMax(root.Right)

	if root.Left == nil && root.Right == nil {
		return root
	}

	var max = root
	if root.Left == nil {
		if max.Val < root.Right.Val {
			max = root.Right
		}
	}
	if root.Right == nil {
		if max.Val < root.Left.Val {
			max = root.Left
		}
	}
	if l != nil && l.Val > max.Val {
		return l
	}
	if r != nil && r.Val > max.Val {
		return r
	}
	return max
}

func checkSwap(root *TreeNode) {
	if root.Left != nil || root.Right != nil {
		if root.Right != nil {
			var min = getMin(root.Right)
			if min.Val < root.Val {
				min.Val, root.Val = root.Val, min.Val
			}
		}

		if root.Left != nil {
			var max = getMax(root.Left)
			if max.Val > root.Val {
				max.Val, root.Val = root.Val, max.Val
			}
		}

		if root.Left != nil && root.Val < root.Left.Val {
			root.Left.Val, root.Val = root.Val, root.Left.Val
		}

		if root.Right != nil && root.Val > root.Right.Val {
			root.Right.Val, root.Val = root.Val, root.Right.Val
		}
	}
}
func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	checkSwap(root)
	recoverTree(root.Left)
	recoverTree(root.Right)
	checkSwap(root)
}
