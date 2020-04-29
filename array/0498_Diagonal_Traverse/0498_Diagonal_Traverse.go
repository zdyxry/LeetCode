package main

import "fmt"

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	n, m := len(matrix), len(matrix[0])
	res := make([]int, n*m)
	var x, y int
	for i := 0; i < len(res); i++ {
		res[i] = matrix[x][y]
		if (x+y)%2 == 0 {
			if y == m-1 {
				x++
			} else if x == 0 {
				y++
			} else {
				y++
				x--
			}
		} else {
			if x == n-1 {
				y++
			} else if y == 0 {
				x++
			} else {
				x++
				y--
			}
		}
	}
	return res
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := findDiagonalOrder(matrix)
	fmt.Println(res)
}
