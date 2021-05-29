func getCharCode(s string) int {
	var i = 0
	for _, v := range s {
		i += int(v)
	}
	return i
}
func groupAnagrams(strs []string) [][]string {
	var m = map[int][]string{}
	for _, v := range strs {
		key := getCharCode(v)
		m[key] = append(m[key], v)
	}
	var test = [][]string{}
	for _, v := range m {
		test = append(test, v)
	}
	return test
}