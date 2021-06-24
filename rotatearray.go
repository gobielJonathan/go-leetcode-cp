func rotate(nums []int, k int) {
	var l = len(nums)
	var c = make([]int, l)

	for i, v := range nums {
		if i+k < l {
			c[i+k] = v
		} else {
			c[(i+k)%l] = v
		}
	}
	for i := 0; i < l; i++ {
		nums[i] = c[i]
	}
}