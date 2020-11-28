package main

import (
	"fmt"
	"strings"
)

func maxRepeating(sequence string, word string) int {
	var w = word
	var count = 0
	for strings.Contains(sequence, word) {
		count++
		word = strings.Repeat(w, count)
	}

	return (len(word) - len(w)) / len(w)
}

func main() {
	sequence := "ababc"
	word := "ab"
	res := maxRepeating(sequence, word)
	fmt.Println(res)
}
