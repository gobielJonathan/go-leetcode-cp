func strStr(haystack string, needle string) int {
	h := len(haystack)
	n := len(needle)

	
	if n == 0 {
		return 0
	}
	if h == 0 {
		return -1
	}
	if h < n {
		return -1
	}

	for i := 0; i < h; i++ {
		if haystack[i] == needle[0] {
			var valid = true
			for y := 0; y < n; y++ {
				if i+y >= h {
					valid = false
					break
				}
				valid = valid && (haystack[i+y] == needle[y])
				if !valid {
					break
				}
			}
			if valid {
				return i
			}
		}
	}
	return -1
}