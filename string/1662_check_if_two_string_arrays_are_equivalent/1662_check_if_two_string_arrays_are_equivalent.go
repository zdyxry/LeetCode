package main

import (
	"fmt"
	"strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	if len(word1) == 0 && len(word2) == 0 {
		return true
	}

	full1 := strings.Join(word1, "")
	full2 := strings.Join(word2, "")
	if len(full1) != len(full2) {
		return false
	}

	for i := 0; i < len(full1); i++ {
		if full1[i] != full2[i] {
			return false
		}
	}
	return true
}

func main() {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}
	res := arrayStringsAreEqual(word1, word2)
	fmt.Println(res)
}
