package main

import (
	"fmt"
	"math"
)

func projectionArea(grid [][]int) int {
	rowNum, colNum := len(grid), len(grid[0])
	xy, xz, yz, rowMax, colMax := 0, 0, 0, make([]int, rowNum), make([]int, colNum)
	for row, v1 := range grid {
		for col, v2 := range v1 {
			if v2 != 0 {
				xy++
			}
			rowMax[row] = int(math.Max(float64(rowMax[row]), float64(grid[row][col])))
			colMax[col] = int(math.Max(float64(colMax[col]), float64(grid[row][col])))
		}
	}
	for _, v := range rowMax {
		xz += v
	}
	for _, v := range colMax {
		yz += v
	}
	return xy + xz + yz
}

func main() {
	grid := [][]int{{2}}
	res := projectionArea(grid)
	fmt.Println(res)
}
