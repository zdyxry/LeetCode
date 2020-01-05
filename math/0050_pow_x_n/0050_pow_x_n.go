package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}

func main() {
	x := float64(4)
	n := 2
	res := myPow(x, n)
	fmt.Println(res)
}
