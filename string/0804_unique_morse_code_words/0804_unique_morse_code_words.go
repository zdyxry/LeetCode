package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	morses := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	transformations := make(map[string]bool)
	for _, s := range words {
		transformation := ""
		runes := []rune(s)
		for _, r := range runes {
			index := r - 'a'
			transformation += morses[index]
		}
		transformations[transformation] = true
	}
	return len(transformations)
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	res := uniqueMorseRepresentations(words)
	fmt.Println(res)
}
