package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	min := math.MaxInt32
	res := [][]int{}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < min {
			min = arr[i] - arr[i-1]
		}
		if min == 1 {
			break
		}
	}
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == min {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}
	return res
}

func main() {
	arr := []int{4, 2, 1, 3}
	res := minimumAbsDifference(arr)
	fmt.Println(res)
}
