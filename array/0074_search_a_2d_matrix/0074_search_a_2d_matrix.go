package main 

import "fmt"

func searchMatrix(mat [][]int, target int) bool {
	m := len(mat)
	if m == 0 {
		return false
	}

	n := len(mat[0])
	if n == 0 {
		return false
	}

	if target < mat[0][0] || mat[m-1][n-1] < target {
		return false
	}

	r := 0
	for r < m && mat[r][0] <= target {
		r++
	}
	r--

	i, j := 0, n-1
	for i <= j {
		med := (i+j)/2
		switch  {
		case mat[r][med] < target:
			i = med +1
		case target < mat[r][med]:
			j = med - 1
		default:
			return true	
		}
	}

	return mat[r][j] == target
}

func main() {
	matrix := [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,50}}
	target := 16

	res := searchMatrix(matrix, target)
	fmt.Println(res)
}