package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	prefix := make([]int, len(arr)+1)
	prefix[0] = 0
	for i := 1; i < len(arr); i++ {
		prefix[i] = prefix[i-1] ^ arr[i-1]
	}
	fmt.Println(prefix)
	res := make([]int, len(queries))
	for i := range queries {
		curr := queries[i]
		res[i] = prefix[curr[1]+1] ^ prefix[curr[0]]
	}
	return res
}

func main() {
	arr := []int{1, 3, 4, 8}
	queries := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}
	res := xorQueries(arr, queries)
	fmt.Println(res)
}
