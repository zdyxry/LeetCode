package main

func countSquares(matrix [][]int) int {
	res := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i > 0 && j > 0 && matrix[i][j] > 0 {
				matrix[i][j] = Min(matrix[i-1][j], Min(matrix[i][j-1], matrix[i-1][j-1])) + 1
			}
			res += matrix[i][j]
		}
	}
	return res
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
