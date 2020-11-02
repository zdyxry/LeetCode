package main

import (
	"fmt"
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	// fmt.Println(points)
	max := -1
	for i := 0; i < len(points)-1; i++ {
		distance := points[i+1][0] - points[i][0]
		if distance > max {
			max = distance
		}
	}

	return max
}

func main() {
	points := [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}
	res := maxWidthOfVerticalArea(points)
	fmt.Println(res)
}
