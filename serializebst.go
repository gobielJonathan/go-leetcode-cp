type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func isLeaf(p *TreeNode) bool {
	if p == nil {
		return false
	}
	return p.Left == nil && p.Right == nil
}

func height(root *TreeNode) int {
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
func takeWithLevel(root *TreeNode, target int, level int) string {
	if root == nil {
		return ""
	}
	var x = takeWithLevel(root.Left, target, level+1)
	if level == target {
		return fmt.Sprintf("%d", root.Val)
	}
	var y = takeWithLevel(root.Right, target, level+1)
	return strings.Join([]string{x, y}, ",")
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var h = height(root)
	var temp = []string{}
	for i := 1; i <= h; i++ {
		var str = takeWithLevel(root, i, 1)
		temp = append(temp, str)
	}
	return strings.Join(temp, ",")
}

func makeTree(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = makeTree(root.Left, val)
	} else if val > root.Val {
		root.Right = makeTree(root.Right, val)
	}
	return root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	var root *TreeNode = nil
	for _, v := range strings.Split(data, ",") {
		v, err := strconv.Atoi(string(v))
		if err == nil {
			root = makeTree(root, v)
		}
	}
	return root
}