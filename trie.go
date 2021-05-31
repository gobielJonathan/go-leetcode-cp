package main

// https://leetcode.com/problems/implement-trie-prefix-tree/

type Trie struct {
	Chars [26]*Trie
	IsEnd bool
}

/** Initialize your data structure here. */
func constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	var curr = this
	for _, v := range word {
		var idx = v - rune('a')
		if curr.Chars[idx] == nil {
			curr.Chars[idx] = new(Trie)
		}
		curr = curr.Chars[idx]
	}
	curr.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	var curr = this
	for _, v := range word {
		var idx = v - rune('a')
		if curr.Chars[idx] == nil {
			return false
		}
		curr = curr.Chars[idx]
	}

	return curr.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	var curr = this
	var isvalid = true
	for _, v := range prefix {
		var idx = v - rune('a')
		isvalid = isvalid && curr.Chars[idx] != nil
		if curr.Chars[idx] == nil || !isvalid {
			return false
		}
		curr = curr.Chars[idx]
	}
	return isvalid
}
