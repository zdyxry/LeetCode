package main

import (
	"fmt"
)

func sortArrayByParity(A []int) []int {
	left := 0
	right := len(A) - 1

	res := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			res[left] = A[i]
			left++
		} else {
			res[right] = A[i]
			right--
		}
	}

	return res

}
func main() {
	A := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(A))
}
