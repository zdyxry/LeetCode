package main

import (
	"fmt"
	"math"
)

func arrangeCoins(n int) int {
	if n <= 0 {
		return 0
	}

	x := math.Sqrt(2*float64(n)+0.25) - 0.5
	return int(x)
}

func main() {
	n := 8
	res := arrangeCoins(n)
	fmt.Println(res)
}
