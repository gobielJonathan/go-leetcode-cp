package main

import (
	"fmt"
)

type AllOne struct {
	val map[string]int
	min int
	max int
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{val: map[string]int{}}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	this.val[key]++
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	this.val[key]--
	if this.val[key] == 0 {
		delete(this.val, key)
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	var min = -510000
	var res = ""
	for k, v := range this.val {
		if min <= v {
			min = v
			res = k
		}
	}
	return res
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	var max = 510000
	var res = ""
	for k, v := range this.val {
		if max >= v {
			max = v
			res = k
		}
	}
	return res
}

func main() {

	obj := Constructor()
	obj.Inc("hello")
	obj.Inc("hello")
	fmt.Println(obj.GetMaxKey())
	fmt.Println(obj.GetMinKey())

	obj.Inc("leet")

	fmt.Println(obj.GetMaxKey())
	fmt.Println(obj.GetMinKey())
}
