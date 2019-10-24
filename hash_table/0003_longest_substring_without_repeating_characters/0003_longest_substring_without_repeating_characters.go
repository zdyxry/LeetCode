package main 

import "fmt"

func lengthOfLongestSubstring(s string) int {
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0

	for i:=0;i<len(s);i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i + 1 - left > maxLen {
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}