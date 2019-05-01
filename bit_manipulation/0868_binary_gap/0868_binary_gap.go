package main

import (
	"fmt"
)

func binaryGap(N int) int {
	res := 0
	gap := 0

	for N > 0 {
		if N&1 == 1 {
			res = max(res, gap)
			gap = 1
		} else if gap > 0 {
			gap++
		}
		N >>= 1
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(binaryGap(5))
}
