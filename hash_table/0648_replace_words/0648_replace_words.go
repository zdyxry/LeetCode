package main 

import (
	"fmt"
	"sort"
	"strings"
)

func replaceWords(dict []string, sentence string) string {
	hasRoot := make(map[string]bool, len(dict))
	hasLen := make([]bool, 101)
	for i := range dict {
		hasRoot[dict[i]] = true
		hasLen[len(dict[i])] = true
	}

	ls := make([]int, 0, 101)
	for i, ok := range hasLen {
		if ok {
			ls = append(ls, i)
		}
	}
	sort.Ints(ls)

	words := strings.Split(sentence, " ")
	for i, w := range words {
		for _, l := range ls {
			if l < len(w) && hasRoot[w[:l]] {
				words[i] = w[:l]
				break
			}
		}
	}

	return strings.Join(words, " ")

}

func main() {
	dict := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	res := replaceWords(dict, sentence)
	fmt.Println(res)
}