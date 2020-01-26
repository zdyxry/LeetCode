package main 

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	lenRow, lenCol := len(grid), len(grid[0])
	res := 0
	for r := 0; r < lenRow; r++ {
		for c := 0; c < lenCol; c++ {
			if grid[r][c] == '1'{
				res++
				dfs(&grid, r, c)
			}
		}
	}
	return res
}

func dfs(grid *[][]byte, r, c int) {
	lenRow, lenCol := len(*grid), len((*grid)[0])
	if r < 0 || r > lenRow || c < 0 || c >= lenCol || (*grid)[r][c] == '0' {
		return
	}
	(*grid)[r][c] = '0'
	dfs(grid, r - 1, c)
	dfs(grid, r + 1, c)
	dfs(grid, r, c - 1)
	dfs(grid, r, c + 1)
}

func main() {
	grid := [][]byte{{'1','1','1','1','0'},{'1','1','0','1','0'},{'1','1','0','0','0'},{'0','0','0','0','0'}}
	res := numIslands(grid)
	fmt.Println(res)
}