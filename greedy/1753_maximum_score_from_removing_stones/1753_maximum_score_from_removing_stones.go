package main

import (
	"fmt"
	"sort"
)

func maximumScore(a int, b int, c int) int {

	stones := []int{a, b, c}
	store := 0
	for {

		sort.Ints(stones)
		if stones[0] == 0 && stones[1] == 0 {
			return store
		}

		if stones[0]+stones[1] <= stones[2] {
			return store + stones[0] + stones[1]
		}

		stones[1], stones[2] = stones[1]-1, stones[2]-1
		store++

	}
	return store
}

func main() {
	a, b, c := 2, 4, 6
	res := maximumScore(a, b, c)
	fmt.Println(res)
}
