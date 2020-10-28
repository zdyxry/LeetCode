package main

import "fmt"

func smallestRangeI(A []int, K int) int {
	min := A[0]
	max := A[0]
	for i := range A {
		if min > A[i] {
			min = A[i]
		}
		if max < A[i] {
			max = A[i]
		}
	}
	if (max - min) >= 2*K {
		return max - min - 2*K
	} else {
		return 0
	}
}

func main() {
	A := []int{1}
	K := 0
	res := smallestRangeI(A, K)
	fmt.Println(res)
}
