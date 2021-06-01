func connectSibling(root *Node, target int, level int) []*Node {
	if root == nil {
		return nil
	}

	var x = connectSibling(root.Left, target, level+1)

	if target == level {
		return []*Node{root}
	}

	var y = connectSibling(root.Right, target, level+1)

	return append(x, y...)
}
func height(root *Node) int {
	if root == nil {
		return 0
	}
	var l = height(root.Left)
	var r = height(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
func connect(root *Node) *Node {
	var h = height(root)

	for i := h; i >= 1; i-- {
		var nodesLevel = connectSibling(root, i, 1)
		var l = len(nodesLevel)
		for j := 0; j < l-1; j++ {
			nodesLevel[j].Next = nodesLevel[j+1]
		}
	}
	return root
}