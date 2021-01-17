package main

import "fmt"

func countGoodRectangles(rectangles [][]int) int {
	maxLen, res := 0, 0
	for _, rectangle := range rectangles {
		curMax := min(rectangle)
		if maxLen < curMax {
			maxLen = curMax
			res = 1
		} else if maxLen == curMax {
			res += 1
		}
	}
	return res
}

func min(arr []int) int {
	if arr[0] <= arr[1] {
		return arr[0]
	}
	return arr[1]
}

func main() {
	rectangles := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}
	res := countGoodRectangles(rectangles)
	fmt.Println(res)
}
