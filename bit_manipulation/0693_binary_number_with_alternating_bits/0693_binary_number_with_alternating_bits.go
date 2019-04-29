package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	prevBit := -1
	for n > 0 {
		lastBit := n & 1
		if lastBit == prevBit {
			return false
		}
		prevBit = lastBit
		n >>= 1
	}

	return true
}

func main() {
	var n int = 7
	fmt.Println(hasAlternatingBits(n))
}
