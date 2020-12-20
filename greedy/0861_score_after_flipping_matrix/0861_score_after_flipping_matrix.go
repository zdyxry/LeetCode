package main

import "math"

func matrixScore(A [][]int) int {

	low := len(A)
	col := len(A[0])

	for i := 0; i < low; i++ {
		if A[i][0] != 1 {
			for j := 0; j < col; j++ {
				if A[i][j] == 0 {
					A[i][j] = 1
				} else {
					A[i][j] = 0
				}
			}
		}
	}
	for i := 0; i < col; i++ {
		res := 0
		for j := 0; j < low; j++ {
			if A[j][i] != 0 {
				res++
			}
		}
		re := int(math.Ceil(float64(low*1.0) / 2))
		if res < re {
			for j := 0; j < low; j++ {
				if A[j][i] != 0 {
					A[j][i] = 0
				} else {
					A[j][i] = 1
				}
			}
		}
	}
	result := 0
	for i := 0; i < low; i++ {
		for j := 0; j < col; j++ {
			if A[i][j] != 0 {
				result += int(math.Pow(float64(2), float64(col-j-1)))
			}

		}
	}

	return result
}
