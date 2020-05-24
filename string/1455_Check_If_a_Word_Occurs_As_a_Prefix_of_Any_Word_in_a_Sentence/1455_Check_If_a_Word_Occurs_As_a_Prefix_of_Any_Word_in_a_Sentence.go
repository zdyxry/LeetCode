package main

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], searchWord) {
			return i + 1
		}
	}
	return -1
}

func main() {
	sentence := "i love eating burger"
	searchWord := "burg"
	res := isPrefixOfWord(sentence, searchWord)
	fmt.Println(res)
}
