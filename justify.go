func fullJustify(words []string, maxWidth int) []string {
	var l = len(words)
	var counter = 0
	var res = []string{}
	var latest = 0

	for i := 0; i < l-1; i++ {
		counter += len(words[i])
		if counter+len(words[i+1]) > maxWidth {
			counter = 0
			var strsplit = words[latest:i]
			var lstrsplit = len(strsplit)
			var space = maxWidth / lstrsplit
			var temp = ""
			for i, v := range strsplit {
				temp += v
				if i < lstrsplit {
					for k := 0; k < space; k++ {
						temp += " "
					}
				}
			}
			res = append(res, temp)
			latest = i
		}
	}
	return res
}	