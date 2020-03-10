package main

import (
	"fmt"
	"math"
)

var target = map[uint8]int{
	'a': 1,
	'e': 2,
	'i': 4,
	'o': 8,
	'u': 16,
}

func findTheLongestSubstring(s string) int {
	var (
		i, curr, rst, st int
		ok               bool
		length           = len(s)
		status           = make([]int, 32)
	)
	for i = 0; i < 32; i++ {
		status[i] = math.MaxInt32
	}
	status[0] = -1
	for i = 0; i < length; i++ {
		if st, ok = target[s[i]]; ok {
			curr ^= st
		}
		if status[curr] == math.MaxInt32 {
			status[curr] = i
		} else {
			rst = max(rst, i-status[curr])
		}
	}
	return rst
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	s := "eleetminicoworoep"
	res := findTheLongestSubstring(s)
	fmt.Println(res)
}
