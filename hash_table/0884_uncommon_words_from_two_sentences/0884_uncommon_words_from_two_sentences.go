package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	words := map[string]bool{}
	ret := []string{}

	var help = func(s string) {
		for _, word := range strings.Split(s, " ") {
			_, ok := words[word]
			words[word] = !ok
		}
	}

	help(A)
	help(B)

	for key, val := range words {
		if val {
			ret = append(ret, key)
		}
	}

	return ret
}

func main() {
	A := "this apple is sweet"
	B := "this apple is sour"

	fmt.Println(uncommonFromSentences(A, B))
}
