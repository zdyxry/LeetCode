package main

import "fmt"

func maxPower(s string) int {
	n, ans := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			n++
		} else {
			n = 1
		}
		if n > ans {
			ans = n
		}
	}
	return ans
}

func main() {
	s := "leetcode"
	res := maxPower(s)
	fmt.Println(res)
}
