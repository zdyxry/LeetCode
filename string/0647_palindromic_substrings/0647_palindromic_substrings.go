package main 

import "fmt"

func countSubstrings(s string) int {
	count := 0
	for i :=0; i<len(s); i++ {
		count += extendPalindrome(s, i, i)
		count += extendPalindrome(s, i, i+1)
	}
	return count
}

func extendPalindrome(s string, left, right int) int {
	res := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		res++
		left--
		right++
	}
	return res
}

func main() {
	s := "abc"
	res := countSubstrings(s)
	fmt.Println(res)
}