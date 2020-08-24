package main

import "fmt"

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	in_degree := make([]int, n)
	size := len(edges)
	ans := make([]int, 0)
	for i := 0; i < size; i++ {
		in_degree[edges[i][1]]++
	}
	for i, j := range in_degree {
		if j == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	n := 6
	edges := [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}
	res := findSmallestSetOfVertices(n, edges)
	fmt.Println(res)
}
