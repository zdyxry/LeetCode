package main

import (
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0{
		return 0
	}

	for i := 1; i < n; i++ {
		for j :=0;j <=i; j++ {
			switch {
			case j == 0:
				triangle[i][0] += triangle[i-1][0]
			case j == i:
				triangle[i][i] += triangle[i-1][i-1]
			case triangle[i-1][j-1] < triangle[i-1][j]:
				triangle[i][j] += triangle[i-1][j-1]
			default:
				triangle[i][j] += triangle[i-1][j]
			}
		}
	}

	min := triangle[n-1][0]
	for j:=1;j < n; j++ {
		if min > triangle[n-1][j]{
			min = triangle[n-1][j]
		}
	}

	return min
}

func main() {
	triangle := [][]int{{2},{3,4}, {6,5,7}, {4,1,8,3}}
	res := minimumTotal(triangle)
	fmt.Println(res)
}