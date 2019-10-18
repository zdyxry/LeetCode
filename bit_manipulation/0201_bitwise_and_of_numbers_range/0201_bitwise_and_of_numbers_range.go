package main 

import "fmt"

func rangeBitwiseAnd(m int, n int) int {
	res := uint(0)

	for m >= 1 && n >= 1{
		if m == n {
			return n << res
		}

		m >>= 1
		n >>= 1
		res++
	}

	return 0
}

func main() {
	m := 26
	n := 30
	res := rangeBitwiseAnd(m, n)
	fmt.Println(res)
}