package main

import "fmt"

func repeatSubstringPattern(s string) bool {
	sz := len(s)
	if sz <= 1 {
		return false
	}

	for i := 2; i <= len(s); i++ {
		if sz%i == 0 {
			n := sz / i
			match := true
			str1 := s[:n]
			for j := n; j < len(s); j += n {
				if s[j:j+n] != str1 {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

func main() {
	s := "aba"
	fmt.Println(repeatSubstringPattern(s))
}
