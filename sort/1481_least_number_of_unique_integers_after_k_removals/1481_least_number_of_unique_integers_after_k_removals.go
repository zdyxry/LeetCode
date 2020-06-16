package main

import (
	"fmt"
	"sort"
)

func findLeastNumOfUniqueInts(arr []int, k int) int {
	cnts := make(map[int]int)
	for _, i := range arr {
		cnts[i]++
	}
	t := make([]int, 0, len(cnts))
	for _, i := range cnts {
		t = append(t, i)
	}
	sort.Ints(t)
	total, idx := 0, 0
	for idx < len(t) && total+t[idx] <= k {
		total += t[idx]
		idx++
	}
	return len(t) - idx
}

func main() {
	arr := []int{5, 5, 4}
	k := 1
	res := findLeastNumOfUniqueInts(arr, k)
	fmt.Println(res)
}
