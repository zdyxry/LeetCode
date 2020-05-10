package main

import "fmt"

func distributeCandies(candies []int) int {
	n := len(candies) / 2
	m := make(map[int]bool, n)
	for _, c := range candies {
		m[c] = true
		if len(m) == n {
			return n
		}
	}
	return len(m)
}

func main() {
	candies := []int{1, 1, 2, 2, 3, 3}
	res := distributeCandies(candies)
	fmt.Println(res)
}
