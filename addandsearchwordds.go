type WordDictionary struct {
	data []string
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	this.data = append(this.data, word)
}

func (this *WordDictionary) Search(word string) bool {
	var isvalid = false
	for _, v := range this.data {
		if len(v) != len(word) {
			continue
		}
		isvalid = true
		for i, y := range word {
			if y == '.' {
				continue
			}
			isvalid = isvalid && y == rune(v[i])
		}
		if isvalid {
			return true
		}
	}
	return isvalid
}