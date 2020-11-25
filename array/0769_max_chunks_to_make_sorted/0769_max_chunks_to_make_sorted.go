package main

import "fmt"

func maxChunksToSorted(arr []int) int {
	if len(arr) < 2 {
		return len(arr)
	}

	count := 0
	cur := arr[0]
	for i, v := range arr {
		cur = max(cur, v)
		if cur == i {
			count++
		}
	}
	return count
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{4, 3, 2, 1, 0}
	res := maxChunksToSorted(arr)
	fmt.Println(res)
}
