package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	var res int // default 0

	row, col := len(grid), len(grid[0])

	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			if grid[x][y] == 1 {
				res = Max(res, dfs(x, y, grid))
			}
		}
	}
	return res
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(x, y int, grid [][]int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0

	count := 1
	count += dfs(x-1, y, grid)
	count += dfs(x+1, y, grid)
	count += dfs(x, y-1, grid)
	count += dfs(x, y+1, grid)
	return count
}

func main() {
	grid := [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}
	res := maxAreaOfIsland(grid)
	fmt.Println(res)
}
