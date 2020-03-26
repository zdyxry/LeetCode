package main

import (
	"fmt"
	sort2 "sort"
)

func diagonalSort(mat [][]int) [][]int {
	var (
		i    int
		lenX = len(mat)
	)
	if lenX == 0 {
		return mat
	}

	var (
		lenY = len(mat[0])
	)

	for i = 0; i < lenX; i++ {
		sort(mat, lenX-1-i, 0, lenX, lenY)
	}
	for i = 1; i < lenY; i++ {
		sort(mat, 0, i, lenX, lenY)
	}
	return mat
}

func sort(mat [][]int, x, y, lenX, lenY int) {
	var (
		tmp        []int
		tmpX, tmpY = x, y
	)
	for tmpX < lenX && tmpY < lenY {
		tmp = append(tmp, mat[tmpX][tmpY])
		tmpX++
		tmpY++
	}
	tmpX = 0
	sort2.Ints(tmp)
	for x < lenX && y < lenY {
		mat[x][y] = tmp[tmpX]
		tmpX++
		x++
		y++
	}
}

func main() {
	mat := [][]int{{3, 3, 1, 1}, {2, 2, 1, 2}, {1, 1, 1, 2}}
	res := diagonalSort(mat)
	fmt.Println(res)
}
