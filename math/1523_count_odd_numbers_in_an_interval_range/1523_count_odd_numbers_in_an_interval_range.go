package main

import "fmt"

func countOdds(low int, high int) int {
	if low+1 == high {
		return 1
	}
	if low == high {
		if low%2 == 0 {
			return 0
		}
		return 1
	}
	if low%2 == 0 {
		low++
	}
	num := (high - low) / 2
	return num + 1
}

func main() {
	low := 3
	high := 7
	res := countOdds(low, high)
	fmt.Println(res)
}
