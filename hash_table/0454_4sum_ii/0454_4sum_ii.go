package main 

import "fmt"

func fourSumCount(A, B, C, D []int) int {
	res := 0
	n := len(A)
	sum := make(map[int]int, n * n)
	for i := 0; i < n; i++ {
		for j := 0;j < n; j++ {
			sum[C[i]+D[j]]++
		}
	}
	fmt.Println(sum)
	for i := 0; i < n; i++ {
		for j:=0;j<n;j++ {
			res += sum[-(A[i]+B[j])]
		}
	}

	return res
}

func main() {
	A := []int{1,2}
	B := []int{-2,-1}
	C := []int{-1,2}
	D := []int{0, 2}
	res := fourSumCount(A,B,C,D)
	fmt.Println(res)
}