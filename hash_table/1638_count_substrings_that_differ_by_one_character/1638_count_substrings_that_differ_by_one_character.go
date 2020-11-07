package main

import "fmt"

func countSubstrings(s string, t string) int {
	n := len(s)
	m := len(t)
	res := 0
	var k, l int
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			diff := 0
			k = i
			l = j
			for k < n && l < m && diff <= 1 {
				if s[k] != t[l] {
					diff++
				}
				if diff == 1 {
					res++
				}
				k++
				l++
			}
		}
	}
	return res
}

func main() {
	s := "aba"
	t := "baba"
	res := countSubstrings(s, t)
	fmt.Println(res)
}
