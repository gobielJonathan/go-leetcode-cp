//blom selesai
func isMatch(s string, p string) bool {
	var i = 0
	var valid = true

	if len(p) == 0 {
		return false
	}

	var isAllMatched = false

	for _, v := range p {
		if v == '*' {
			isAllMatched = true
			continue
		}
		if v == '?' {
			valid = valid && true
			i++
			continue
		}
		if isAllMatched {
			var l = len(s)
			for h := 0; h < l; h++ {
				if rune(s[h]) == v {
					isAllMatched = false
					i = h
					break
				}
			}
		}
		valid = valid && rune(s[i]) == v
		if !valid {
			return false
		}
		i++
		if i > len(s)-1 {
			break
		}
	}
	if isAllMatched {
		return true
	}
	if i < len(s) {
		return false
	}
	if valid {
		if len(p) > len(s) {
			return false
		}
	}
	return valid
}