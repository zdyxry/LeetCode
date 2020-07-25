package main

import "fmt"

func shiftGrid(grid [][]int, k int) [][]int {
	res := make([]int, 0, 10)
	lenth := len(grid)
	w := len(grid[0])
	for i := 0; i < lenth; i++ {
		res = append(res, grid[i]...)
	}
	k = k % (lenth * w)
	res = append(res[lenth*w-k:], res[:lenth*w-k]...)
	tmp := make([][]int, 0, 10)
	for i := 0; i < lenth; i++ {
		tmp = append(tmp, res[i*w:(i+1)*w])
	}
	return tmp
}

func main() {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := shiftGrid(grid, 1)
	fmt.Println(res)
}
