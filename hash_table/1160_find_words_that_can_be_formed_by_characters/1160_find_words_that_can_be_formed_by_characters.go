package main

import "fmt"

func countCharacters(words []string, chars string) int {
	magic := [26]int{}
	for _, c := range chars {
		magic[c-'a']++
	}
	good := func(word string, magic [26]int) bool {
		for _, c := range word {
			magic[c-'a']--
			if magic[c-'a'] < 0 {
				return false
			}
		}
		return true
	}

	goodWord := 0
	for _, word := range words {
		if good(word, magic) {
			goodWord += len(word)
		}
	}
	return goodWord
}

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	res := countCharacters(words, chars)
	fmt.Println(res)
}
