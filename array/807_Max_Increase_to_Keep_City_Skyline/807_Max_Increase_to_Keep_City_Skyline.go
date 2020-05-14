package main

import (
	"fmt"
	"math"
)

func maxIncreaseKeepingSkyline(grid [][]int) int {
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return 0
	}

	maxRow, maxCol := make([]int, len(grid)), make([]int, len(grid[0]))
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		max := 0
		for j := 0; j < cols; j++ {
			max = int(math.Max(float64(max), float64(grid[i][j])))
		}
		maxRow[i] = max
	}
	for i := 0; i < cols; i++ {
		max := 0
		for j := 0; j < rows; j++ {
			max = int(math.Max(float64(max), float64(grid[j][i])))
		}
		maxCol[i] = max
	}
	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res += (int(math.Min(float64(maxRow[i]), float64(maxCol[j]))) - grid[i][j])
		}
	}
	return res
}

func main() {
	grid := [][]int{{3, 0, 8, 4}, {2, 4, 5, 7}, {9, 2, 6, 3}, {0, 3, 1, 0}}
	res := maxIncreaseKeepingSkyline(grid)
	fmt.Println(res)
}
