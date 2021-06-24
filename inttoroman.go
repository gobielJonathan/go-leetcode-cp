func intToRoman(num int) string {
	var m = map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	var keys = []int{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	var target = num
	var str = ""

	for _, v := range keys {
		if v <= target {
			str += m[v]
			target -= v
			for target >= v {
				str += m[v]
				target -= v
			}
		}
	}

	return str
}
