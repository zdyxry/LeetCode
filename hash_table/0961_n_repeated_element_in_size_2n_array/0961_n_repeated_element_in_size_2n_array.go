package main

import "fmt"

func repeatedNTimes(A []int) int {
	for i := 0; i < len(A)-2; i++ {
		if A[i] == A[i+1] || A[i] == A[i+2] {
			return A[i]
		}
	}

	return A[len(A)-1]
}

func main() {
	A := []int{1, 2, 3, 3}
	res := repeatedNTimes(A)
	fmt.Println(res)
}
