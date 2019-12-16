package main 

import (
	"fmt"
	"math"
)

func nthSuperUglyNumber(n int, primes []int) int {
	ugly := make([]int, n)
	idx := make([]int, len(primes))
	val := make([]int, len(primes))
	for i := 0; i < len(primes); i++ {
		val[i] = 1
	}

	next := 1
	for i := 0; i < n ; i++ {
		ugly[i] = next
		next = math.MaxInt32
		for j := 0; j <len(primes); j++ {
			if val[j] == ugly[i] {
				val[j] = ugly[idx[j]] * primes[j]
				idx[j]++
			}
			next = min(next, val[j])
		}
	}
	return ugly[n-1]
}

func min(x,y int) int {
	if x < y {
		return x
	} 
	return y
}

func main() {
	n := 12
	primes := []int{2,7, 13, 19}
	res := nthSuperUglyNumber(n, primes)
	fmt.Println(res)
}

