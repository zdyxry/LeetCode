package  main 

import (
	"fmt"
	"sort"
)

func largestPerimeter(A []int) int {
	sort.Ints(A)

	for i := len(A) - 1; i >= 0 ; i-- {
		if i - 2 < 0 {
			return 0
		}
		if A[i] < A[i-1]+A[i-2] {
			return A[i] + A[i-1] + A[i-2]
		}
	}

	return 0
}

func main() {
	A := []int{1,2,1}
	res := largestPerimeter(A)
	fmt.Println(res)
}