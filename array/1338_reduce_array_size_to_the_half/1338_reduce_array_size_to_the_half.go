package main

import (
	"fmt"
	"sort"
)

func minSetSize(arr []int) int {
	m := make(map[int]int)
	var r []int
	for _, v := range arr {
		m[v]++
	}

	for _, v := range m {
		r = append(r, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(r)))

	var count int
	for idx, v := range r {
		count += v
		if count > len(arr)/2 {
			return idx + 1
		}
	}
	return -1
}

func main() {
	arr := []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}
	res := minSetSize(arr)
	fmt.Println(res)
}
