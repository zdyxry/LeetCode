package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {
	indexes := [26]int{}
	for _, v := range "qwertyuiop" {
		indexes[v-'a'] = 1
	}

	for _, v := range "asdfghjkl" {
		indexes[v-'a'] = 2
	}

	for _, v := range "zxcvbnm" {
		indexes[v-'a'] = 3
	}
	res := make([]string, 0)
	for _, word := range words {
		temp := strings.ToLower(word)
		found := true
		for _, c := range temp {
			if indexes[c-'a'] != indexes[temp[0]-'a'] {
				found = false
				break
			}
		}

		if found {
			res = append(res, word)
		}
	}
	return res

}

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	res := findWords(words)
	fmt.Println(res)
}
