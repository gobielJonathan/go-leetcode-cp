
func isPalindrome(x int) interface{} {
	if x < 0 {
		return false
	}
	if x >= 0 && x <= 9 {
		return true
	}

	len := int(math.Log10(float64(x)) + 1)

	var valid = true

	var p = int(math.Pow(10, float64(len-1)))
	var y = 1

	for i := len; i >= len/2; i-- {

		valid = valid && (x/p%10 == x/y%10)

		if p == 0 {
			p = 10
		}
		p = p / 10
		y = y * 10
	}

	return valid
}
