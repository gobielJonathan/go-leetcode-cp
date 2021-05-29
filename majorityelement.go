func majorityElement(nums []int) []int {
	var temp = map[int]int{}

	for _, v := range nums {
		temp[v]++
	}

	var arr = []int{}
	var min = len(nums) / 3
	for key, v := range temp {
		if v > min {
			arr = append(arr, key)
		}
	}
	return arr
}

func majorityElement(nums []int) int {
	var temp = map[int]int{}

	for _, v := range nums {
		temp[v]++
	}

	var min = len(nums) / 2
	for key, v := range temp {
		if v > min {
			return key
		}
	}
	return 0
}