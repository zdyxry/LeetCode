package main

import "fmt"

func maxUniqueSplit(s string) int {
	return helper(s, map[string]bool{})
}

func helper(s string, seen map[string]bool) int {
	ans := 0
	for i := 1; i <= len(s); i++ {
		if !seen[s[0:i]] {
			seen[s[0:i]] = true
			ans = max(ans, 1+helper(s[i:], seen))
			seen[s[0:i]] = false
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "ababccc"
	res := maxUniqueSplit(s)
	fmt.Println(res)
}
