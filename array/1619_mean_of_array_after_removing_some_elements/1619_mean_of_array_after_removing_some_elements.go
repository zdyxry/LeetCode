package main

import (
	"fmt"
	"sort"
)

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := int(len(arr) * 5 / 100)
	sum := 0
	cnt := 0
	for idx, v := range arr {
		if n <= idx && idx < len(arr)-n {
			sum += v
			cnt++
		}
	}
	return float64(sum) / float64(cnt)
}

func main() {
	arr := []int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4}
	res := trimMean(arr)
	fmt.Println(res)
}
