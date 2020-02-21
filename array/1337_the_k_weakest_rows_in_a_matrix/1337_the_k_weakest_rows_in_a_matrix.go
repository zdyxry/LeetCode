package main

import (
	"fmt"
	"sort"
)

type line struct {
	row   int
	count int
}

func kWeakestRows(mat [][]int, k int) []int {
	strenght := make([]line, len(mat))

	for i := 0; i < len(mat); i++ {
		strenght[i].row = i
		for j := 0; j < len(mat[i]); j++ {
			strenght[i].count += mat[i][j]
		}
	}

	sort.Slice(strenght, func(i, j int) bool {
		if strenght[i].count == strenght[j].count {
			return strenght[i].row < strenght[j].row
		}

		return strenght[i].count < strenght[j].count
	})

	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = strenght[i].row
	}

	return ans
}

func main() {
	mat := [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}
	res := kWeakestRows(mat, 3)
	fmt.Println(res)
}
