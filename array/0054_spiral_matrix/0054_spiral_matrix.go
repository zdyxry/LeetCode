package main 

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}

	m, n := len(matrix), len(matrix[0])

	next := nextFunc(m, n)

	res := make([]int, m*n)
	for i := range res {
		x,y := next()
		res[i] = matrix[x][y]
	}

	return res
}

func nextFunc(m,n int) func() (int, int) {
	top, down := 0, m -1
	left, right := 0, n - 1
	x, y := 0, -1
	dx, dy := 0, 1
	return func() (int, int) {
		x += dx
		y += dy
		switch {
		case y+dy > right:
			top++
			dx, dy = 1, 0
		case x+dx > down:
			right--
			dx, dy = 0, -1
		case y + dy < left:
			down--
			dx, dy = -1,0
		case x + dx < top:
			left++
			dx, dy = 0, 1
		}
		return x,y
	}
}

func main() {
	//[[1,2,3],[4,5,6],[7,8,9]]
	matrix := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	res := spiralOrder(matrix)
	fmt.Println(res)
}