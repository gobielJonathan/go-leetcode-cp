func replaceWords(dictionary []string, sentence string) string {

	var res = []string{}

	for _, v := range strings.Split(sentence, " ") {
		var isFound = false
		var temp = ""
		for _, d := range dictionary {
			if strings.Index(v, d) >= 0 {
				isFound = true
				temp = d
				break
			}
		}
		if isFound {
			res = append(res, temp)
			continue
		}
		res = append(res, v)
	}
	return strings.Join(res, " ")
}