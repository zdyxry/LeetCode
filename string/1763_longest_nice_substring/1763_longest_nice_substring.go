package main

import "fmt"

func longestNiceSubstring(s string) string {
	check := func(s string) bool {
		counts := map[rune]int{}
		for _, ch := range s {
			if ch < 'Z' {
				ch += 'a' - 'A'
				counts[ch] |= 1
			} else {
				counts[ch] |= 2
			}
		}

		for _, count := range counts {
			if count != 3 {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(s)-1; i++ {
		for j := 0; j <= i; j++ {
			sub := s[j : len(s)-i+j]

			if check(sub) {
				return sub
			}
		}
	}
	return ""
}

func main() {
	s := "YazaAay"
	res := longestNiceSubstring(s)
	fmt.Println(res)
}
