package main

import "fmt"

func maxVowels(s string, k int) int {
	res := 0
	dp := make([]int, len(s)+1)
	count := 0
	for i := 0; i < len(s); i++ {
		dp[i] = count
		if check(s[i]) == true {
			count++
		}
	}
	dp[len(s)] = count
	for i := k; i <= len(s); i++ {
		if dp[i]-dp[i-k] > res {
			res = dp[i] - dp[i-k]
		}
	}
	return res
}

func check(str byte) bool {
	if str == 'a' || str == 'e' || str == 'i' || str == 'o' || str == 'u' {
		return true
	}
	return false
}

func main() {
	s := "abciiidef"
	k := 3
	res := maxVowels(s, k)
	fmt.Println(res)
}
