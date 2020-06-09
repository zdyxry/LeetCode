package main

import (
	"fmt"
	"sort"
)

func getStrongest(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}

	res := make([]int, 0, k)
	sort.IntSlice(arr).Sort()

	mIdx := ((len(arr) - 1) / 2)
	m := arr[mIdx]

	low, high := 0, len(arr)-1
	for low <= high && len(res) < k {
		if low == high {
			res = append(res, m)
			break
		}

		if (m - arr[low]) > (arr[high] - m) {
			res = append(res, arr[low])
			low++
		} else {
			res = append(res, arr[high])
			high--
		}
	}

	return res
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	k := 2
	res := getStrongest(arr, k)
	fmt.Println(res)
}
