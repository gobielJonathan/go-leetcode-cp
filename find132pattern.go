func find132pattern(nums []int) bool {
	var l = len(nums)
	var isValid = false
	for i := 0; i < l-2; i++ {
		var temp = nums[i+1] > nums[i+2] &&
			nums[i+1] > nums[i] &&
			nums[i+2] > nums[i]

		isValid = isValid || temp
		if isValid {
			return true
		}
	}
	return isValid
}