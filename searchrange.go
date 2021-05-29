
func searchRange(nums []int, target int) []int {
	var start = -1
	var end = -1
	l := len(nums)
	for i := 0; i < l; i++ {
		var r = l - i - 1
		if start == -1 && nums[i] == target {
			start = i
		}
		if end == -1 && nums[r] == target {
			end = r
		}

		if start != -1 && end != -1 {
			break
		}
	}
	return []int{start, end}
}
