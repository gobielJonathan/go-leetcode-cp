type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	var newNode = new(Node)
	newNode.Val = node.Val
	neighbors := []*Node{}
	for _, v := range node.Neighbors {
		neighbors = append(neighbors, cloneGraph(v))
	}
	newNode.Neighbors = neighbors
	return newNode
}