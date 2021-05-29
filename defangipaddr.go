func defangIPaddr(address string) string {
	var str = ""
	for _, v := range address {
		if v == '.' {
			str += "[.]"
			continue
		}
		str += string(v)
	}
	return str
}