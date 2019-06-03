package main

import (
	"fmt"
)

func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])

	for i := 0; i+1 < m; i++ {
		for j := 0; j+1 < n; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}

	return true
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	fmt.Println(isToeplitzMatrix(matrix))
}
