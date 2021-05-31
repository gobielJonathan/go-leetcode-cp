type MagicDictionary struct {
	Data []string
}

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.Data = append(this.Data, dictionary...)
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for _, v := range this.Data {
		if len(v) != len(searchWord) {
			continue
		}
		var count = 0
		for i, y := range v {
			if string(searchWord[i]) != string(y) {
				count++
			}
		}
		if count == 1 {
			return true
		}
	}
	return false
}