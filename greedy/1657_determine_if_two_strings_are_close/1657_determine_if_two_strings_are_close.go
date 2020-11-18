package main

import "fmt"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	maps := make(map[rune]int)

	for _, val := range word1 {
		maps[val]++
	}

	for _, val := range word2 {
		if maps[val] == 0 {
			return false
		}
		maps[val+'a']++
	}

	s := 0
	for k, v := range maps {
		if k <= 'z' {
			if maps[k+'a'] == 0 {
				return false
			}
		}
		s ^= v
	}
	return s == 0
}

func main() {
	word1 := "abc"
	word2 := "bca"
	res := closeStrings(word1, word2)
	fmt.Println(res)
}
