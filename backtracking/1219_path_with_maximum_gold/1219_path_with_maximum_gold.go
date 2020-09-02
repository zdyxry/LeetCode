package main

import "fmt"

var m, n int

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func dfs(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}

	result := 0
	temp := grid[i][j]
	grid[i][j] = 0

	if i > 0 {
		result = max(result, temp+dfs(grid, i-1, j))
	}

	if j > 0 {
		result = max(result, temp+dfs(grid, i, j-1))
	}

	if i < m-1 {
		result = max(result, temp+dfs(grid, i+1, j))
	}

	if j < n-1 {
		result = max(result, temp+dfs(grid, i, j+1))
	}

	grid[i][j] = temp

	return result
}

func getMaximumGold(grid [][]int) int {
	result := 0
	m = len(grid)
	n = len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result = max(result, dfs(grid, i, j))
		}
	}

	return result
}

func main() {
	grid := [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}
	res := getMaximumGold(grid)
	fmt.Println(res)
}
