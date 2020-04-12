package main

import (
	"fmt"
	"sort"
	"strings"
)

type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func stringMatching(words []string) []string {
	sort.Sort(ByLen(words))
	res := []string{}
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) {
				res = append(res, words[i])
			}
		}
	}

	unique := []string{}
	for _, v := range res {
		skip := false
		for _, u := range unique {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, v)
		}
	}
	return unique
}

func main() {
	words := []string{"mass", "as", "hero", "superhero"}
	res := stringMatching(words)
	fmt.Println(res)
}
