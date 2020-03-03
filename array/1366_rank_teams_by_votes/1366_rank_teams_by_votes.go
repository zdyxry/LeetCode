package main

import (
	"fmt"
	"sort"
)

func rankTeams(votes []string) string {
	bs := make([]byte, 26)
	for i := range bs {
		bs[i] = byte(i + 'A')
	}
	fmt.Println(bs)
	cnts := ['Z' + 1][26]int{}
	for _, s := range votes {
		for i, c := range s {
			cnts[c][i]++
		}
	}
	sort.Slice(bs, func(i, j int) bool {
		a, b := bs[i], bs[j]
		ca, cb := cnts[a], cnts[b]
		for k := range ca {
			if ca[k] != cb[k] {
				return ca[k] > cb[k]
			}
		}
		return a < b
	})

	return string(bs[:len(votes[0])])
}

func main() {
	votes := []string{"ABC", "ACB", "ABC", "ACB", "ACB"}
	res := rankTeams(votes)
	fmt.Println(res)
}
