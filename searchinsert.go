func searchInsert(nums []int, target int) int {
	var l = 0
	var r = len(nums)
	if nums[l] >= target {
		return 0
	}
	if nums[r-1] == target {
		return r - 1
	}
	if nums[r-1] < target {
		return r
	}
	for l < r {
		var mid = int(l + (r-l)/2)
		if mid == 0 {
			for i := 0; i < len(nums); i++ {
				if nums[i] > target {
					return i
				}
			}
		}
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r
}