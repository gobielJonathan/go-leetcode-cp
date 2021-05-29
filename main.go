package main

import "fmt"

func singleNumber(nums []int) int {
	var m = map[int]int{}
	for _, v := range nums {
		m[v]++
	}
	var singlest = -1
	for _, v := range m {
		d := m[singlest]
		c := m[v]
		if singlest == -1 {
			singlest = v
		}
		if d > c {
			singlest = v
		}
	}
	return singlest
}

func main() {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))

}
