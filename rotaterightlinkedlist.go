func rotateRight(head *ListNode, k int) *ListNode {

	var nums = []int{}
	var curr = head
	for curr != nil {
		nums = append(nums, curr.Val)
		curr = curr.Next
	}
	var l = len(nums)
	var c = make([]int, l)
	for i, v := range nums {
		if i+k < l {
			c[i+k] = v
		} else {
			c[(i+k)%l] = v
		}
	}

	var res *ListNode = nil
	for _, v := range c {
		var n = ListNode{Val: v}
		if res == nil {
			res = &n
		} else {
			var curr = res
			for curr.Next != nil {
				curr = curr.Next
			}
			curr.Next = &n
		}
	}
	return res
}

