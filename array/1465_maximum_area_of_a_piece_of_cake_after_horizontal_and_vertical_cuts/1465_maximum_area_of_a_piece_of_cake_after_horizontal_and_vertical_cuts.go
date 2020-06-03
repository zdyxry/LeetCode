package main 

import (
	"fmt"
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
    maxH, maxV := 0, 0
	pre := 0
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	for i := 0; i < len(horizontalCuts); i++ {
		maxH = max(maxH, horizontalCuts[i] - pre)
		pre = horizontalCuts[i]
	}
	maxH = max(maxH, h - pre)
	pre = 0
	for i := 0; i < len(verticalCuts); i++ {
		maxV = max(maxV, verticalCuts[i] - pre)
		pre = verticalCuts[i]
	}
	maxV = max(maxV, w - pre)
	return (maxH * maxV) % 1000000007
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	h := 5
	w := 4
	hor := []int{1,2,4}
	ver := []int{1,3}
	res := maxArea(h,w,hor,ver)
	fmt.Println(res)
}