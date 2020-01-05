package main

import (
	"fmt"
	"math"
)

func divide(m, n int) int {
	if n == 0 {
		return math.MaxInt32
	}

	signM, absM := analysis(m)
	signN, absN := analysis(n)

	res, _ := d(absM, absN, 1)
	if signM != signN {
		res = res - res - res
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}

func analysis(num int) (sign, abs int) {
	sign = 1
	abs = num
	if num < 0 {
		sign = -1
		abs = num - num - num
	}

	return
}

func d(m, n, count int) (res, remainder int) {
	switch {
	case m < n:
		return 0, m
	case n <= m && m < n+n:
		return count, m - n
	default:
		res, remainder = d(m, n+n, count+count)
		if remainder >= n {
			return res + count, remainder - n
		}
		return res, remainder
	}
}

func main() {
	m := 10
	n := 3
	result := divide(m, n)
	fmt.Println(result)
}
