func romanToInt(s string) int {
	var romanMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	l := len(s)
	count := 0

	for i := 0; i < l; i++ {
		key := string(s[i])
		p := romanMap[key]
		count += p

		fmt.Println(count)

		if i < 1 {
			continue
		}

		if key == "V" && string(s[i-1]) == "I" {
			s := romanMap["I"]
			y := romanMap["V"]
			z := count - (y - s)
			count -= z
		}
		if key == "X" && string(s[i-1]) == "I" {
			s := romanMap["I"]
			y := romanMap["X"]
			z := count - (y - s)
			count -= z
		}
		if key == "L" && string(s[i-1]) == "X" {
			s := romanMap["X"]
			y := romanMap["L"]
			z := count - (y - s)
			count -= z
		}
		if key == "C" && string(s[i-1]) == "X" {
			s := romanMap["X"]
			y := romanMap["C"]
			z := count - (y - s)
			count -= z
		}
		if key == "D" && string(s[i-1]) == "C" {
			s := romanMap["C"]
			y := romanMap["D"]
			z := count - (y - s)
			count -= z
		}
		if key == "M" && string(s[i-1]) == "C" {
			s := romanMap["C"]
			y := romanMap["M"]
			z := count - (y - s)
			count -= z
		}
	}
	return count
}
	