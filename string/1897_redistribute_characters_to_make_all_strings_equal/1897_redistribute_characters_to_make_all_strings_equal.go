package main

import "fmt"

func makeEqual(words []string) bool {
	cnt := make(map[rune]int)
	for _, word := range words {
		for _, i := range word {
			cnt[i] += 1
		}
	}
	l := len(words)
	for _, v := range cnt {
		if v%l != 0 {
			return false
		}
	}
	return true

}

func main() {
	//["abc","aabc","bc"]
	words := []string{"abc", "aabc", "bc"}
	result := makeEqual(words)
	fmt.Println(result)
}
