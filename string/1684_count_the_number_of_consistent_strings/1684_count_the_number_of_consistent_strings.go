package main

import "fmt"

func countConsistentStrings(allowed string, words []string) int {
	res := len(words)
	allowedChars := [26]bool{}
	for _, c := range allowed {
		allowedChars[c-'a'] = true
	}
	for _, w := range words {
		for _, c := range w {
			if !allowedChars[c-'a'] {
				res--
				break
			}
		}
	}

	return res
}

func main() {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	res := countConsistentStrings(allowed, words)
	fmt.Println(res)
}
