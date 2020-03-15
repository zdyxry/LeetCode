package main

import "fmt"

func luckyNumbers(matrix [][]int) []int {
	rows, cols := len(matrix), len(matrix[0])
	res := make([]int, 0)
	for i := 0; i < rows; i++ {
		left, right := 0, cols-1
		for left < right {
			if matrix[i][left] > matrix[i][right] {
				left++
			} else {
				right--
			}
		}
		up, down := 0, rows-1
		for up < down {
			if matrix[up][left] > matrix[down][left] {
				down--

			} else {
				up++
			}
		}
		if i == up {
			res = append(res, matrix[up][left])
		}
	}
	return res
}

func main() {
	matrix := [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}
	res := luckyNumbers(matrix)
	fmt.Println(res)
}
