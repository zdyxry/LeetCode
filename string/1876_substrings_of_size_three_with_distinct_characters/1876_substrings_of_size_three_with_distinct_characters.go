package main

import "fmt"

func countGoodSubstrings(s string) int {
	cnt := 0
	for i := 0; i <= len(s)-3; i++ {
		if s[i] != s[i+1] && s[i+1] != s[i+2] && s[i+2] != s[i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	s := "xyzzaz"
	res := countGoodSubstrings(s)
	fmt.Println(res)
}
