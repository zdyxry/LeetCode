package main

import (
	"fmt"
	"sort"
)

func maxCoins(piles []int) int {
	sort.Ints(piles)

	ans := 0
	for n := 1; n <= len(piles)/3; n++ {
		ans += piles[len(piles)-2*n]
	}
	return ans
}

func main() {
	piles := []int{1, 2, 4, 2, 7, 8}
	res := maxCoins(piles)
	fmt.Println(res)
}
