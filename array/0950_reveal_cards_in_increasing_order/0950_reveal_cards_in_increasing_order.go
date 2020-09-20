package main

import (
	"fmt"
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	sort.Sort(sort.IntSlice(deck))
	out := []int{}
	for len(deck) > 0 {
		out = append(deck[len(deck)-1:len(deck)], out...)
		if len(deck) > 1 {
			deck = deck[0 : len(deck)-1]
			out = append(out[len(out)-1:len(out)], out[0:len(out)-1]...)
		} else {
			break
		}
	}
	return out
}

func main() {
	deck := []int{17, 13, 11, 2, 3, 5, 7}
	res := deckRevealedIncreasing(deck)
	fmt.Println(res)
}
