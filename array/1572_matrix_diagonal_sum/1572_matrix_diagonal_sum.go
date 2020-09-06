package main

import "fmt"

func diagonalSum(mat [][]int) int {
	ans := 0
	for i := 0; i < len(mat); i++ {
		ans += mat[i][i]
		if i != len(mat)-1-i {
			ans += mat[i][len(mat)-1-i]
		}
	}
	return ans
}

func main() {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := diagonalSum(mat)
	fmt.Println(res)
}
