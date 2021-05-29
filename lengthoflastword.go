func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	l := len(s)

	count := 0

	for i := l - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		}
		count++
	}
	return count
}