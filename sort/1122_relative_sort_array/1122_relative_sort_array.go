package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	counter := map[int]int{}
	others := []int{}
	dest := make([]int, 0, len(arr2))
	for _, v := range arr2 {
		counter[v] = 0
	}
	for _, v := range arr1 {
		if _, exists := counter[v]; exists {
			counter[v]++
		} else {
			others = append(others, v)
		}
	}
	for _, v := range arr2 {
		for i := 0; i < counter[v]; i++ {
			dest = append(dest, v)
		}
	}
	sort.Ints(others)
	dest = append(dest, others...)
	return dest
}

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	res := relativeSortArray(arr1, arr2)
	fmt.Println(res)
}
