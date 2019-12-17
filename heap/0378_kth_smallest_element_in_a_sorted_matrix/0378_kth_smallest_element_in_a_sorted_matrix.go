package main 

import (
	"fmt"
)

func kthSmallest(mat [][]int, k int) int {
	n := len(mat)

	lo, hi := mat[0][0], mat[n-1][n-1]

	for lo < hi {
		mid := lo + (hi-lo) >> 2
		count := 0
		j := n-1
		for i := 0; i < n; i++ {
			for j >= 0 && mat[i][j] > mid {
				j--
			}
			count += j + 1
		}

		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	mat := [][]int{{1,5,9},{10,11,13},{12,13,15}}
	k := 8
	res := kthSmallest(mat, k)
	fmt.Println(res)
}