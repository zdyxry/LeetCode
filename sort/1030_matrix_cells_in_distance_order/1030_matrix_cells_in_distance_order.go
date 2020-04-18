package main

import "fmt"

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	m := make(map[int][][]int, 0)
	maxDist := 0
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			dist := abs(r-r0) + abs(c-c0)
			if dist > maxDist {
				maxDist = dist
			}
			m[dist] = append(m[dist], []int{r, c})
		}
	}
	fmt.Println(m)
	result := make([][]int, R*C)
	k := 0
	for i := 0; i <= maxDist; i++ {
		for _, v := range m[i] {
			result[k] = v
			k++
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	R, C := 1, 2
	r0, c0 := 0, 0
	res := allCellsDistOrder(R, C, r0, c0)
	fmt.Println(res)
}
