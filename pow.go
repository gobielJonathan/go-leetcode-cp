func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	fmt.Println(x, n)
	c := x
	for i := 1; i < n; i++ {
		c *= x
	}
	return c
}
