package main

import "fmt"

func numSpecial(mat [][]int) int {
	m, n, ans := len(mat), len(mat[0]), 0
	row, col := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				row[i]++
				col[j]++
			}
		}
	}

	pool := []int{}
	for j := 0; j < n; j++ {
		if col[j] == 1 {
			pool = append(pool, j)
		}
	}

	for i := 0; i < m; i++ {
		if row[i] != 1 {
			continue
		}
		for j := range pool {
			if mat[i][pool[j]] == 1 {
				ans++
				break
			}
		}
	}
	return ans
}

func main() {
	mat := [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}
	res := numSpecial(mat)
	fmt.Println(res)
}
