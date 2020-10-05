package main

import "fmt"

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	matrix := make([][]int, 0, len(rowSum))

	for i := 0; i < len(rowSum); i++ {
		matrix = append(matrix, make([]int, len(colSum)))
		for j := 0; j < len(colSum); j++ {
			val := min(rowSum[i], colSum[j])
			matrix[i][j] = val
			rowSum[i] = rowSum[i] - val
			colSum[j] = colSum[j] - val
		}
	}

	return matrix
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	rowSum := []int{3, 8}
	colSum := []int{4, 7}
	res := restoreMatrix(rowSum, colSum)
	fmt.Println(res)
}
