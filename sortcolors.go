func sortColors(nums []int) {
	var l = len(nums)
	for i := 0; i < l; i++ {
		var sorted = true
		for j := 0; j < l-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				sorted = false
			}
		}
		if sorted {
			break
		}
	}
}