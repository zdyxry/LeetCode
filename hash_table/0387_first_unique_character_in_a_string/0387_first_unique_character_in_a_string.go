package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	rec := make([]int, 26)
	for _, b := range s {
		rec[b-'a']++
	}

	for i, b := range s {
		if rec[b-'a'] == 1 {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
