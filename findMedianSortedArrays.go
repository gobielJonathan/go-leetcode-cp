func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var temp = append(nums1, nums2...)
	sort.Ints(temp)
	var l = len(temp)
	if l == 0 {
		return float64(0)
	}
	if l%2 == 1 {
		return float64(temp[l/2])
	}

	return (float64(temp[l/2]) + float64(temp[l/2-1])) / float64(2)
}