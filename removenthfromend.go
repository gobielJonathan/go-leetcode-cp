func getN(h *ListNode) int {
	if h == nil {
		return 0
	}
	return 1 + getN(h.Next)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var idx = getN(head)

	head.Next = removeNthFromEnd(head.Next, n)
	if idx == n {
		return head.Next
	}
	return head
}