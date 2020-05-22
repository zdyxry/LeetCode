package main

import "fmt"

func isMonotonic(A []int) bool {
	if len(A) <= 2 {
		return true
	}
	noInc, noDec := true, true
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			noDec = false
		}
		if A[i] < A[i+1] {
			noInc = false
		}

	}
	return noDec || noInc
}

func main() {
	A := []int{1, 2, 2, 3}
	res := isMonotonic(A)
	fmt.Println(res)
}
