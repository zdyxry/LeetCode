package main

import (
	"fmt"
	"math"
)

func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	l := int(math.Pow(2, float64(n)) - 1)
	mid := int(l/2 + 1)
	if k == mid {
		return '1'
	} else if k < mid {
		return findKthBit(n-1, k)
	} else {
		return byte(1 - int(findKthBit(n-1, int(l-k+1))))
	}

}

func main() {
	n := 3
	k := 1
	res := findKthBit(n, k)
	fmt.Println(string(res))
}
