package main 

import (
	"fmt"
)

func findAndReplacePattern(words []string, pattern string) []string {
	np := normalize(pattern)
	res := make([]string, 0, len(words))
	for _, w := range words {
		if normalize(w) == np {
			res = append(res, w)
		}
	}
	return res
}

func normalize(str string) string {
	m := make(map[rune]byte, len(str))

	for _, c := range str {
		if _, ok := m[c]; !ok {
			m[c] = byte(len(m)) + 'a'
		}
	}
	res := make([]byte, len(str))
	for i, c := range str {
		res[i] = m[c]
	}
	return string(res)
}

func main() {
	words := []string{"abc","deq","mee","aqq","dkd","ccc"}
	pattern := "abb"
	res := findAndReplacePattern(words, pattern)
	fmt.Println(res)
}