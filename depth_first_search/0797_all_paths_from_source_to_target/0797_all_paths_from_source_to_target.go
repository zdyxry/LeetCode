package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	var ans [][]int
	helper(graph, []int{0}, &ans)
	return ans
}

func helper(graph [][]int, path []int, ans *[][]int) {
	pos := path[len(path)-1]
	target := len(graph) - 1
	for _, child := range graph[pos] {
		if child == target {
			cp := make([]int, len(path))
			copy(cp, path)
			cp = append(cp, child)
			*ans = append(*ans, cp)
		} else {
			helper(graph, append(path, child), ans)
		}
	}
	return
}

func main() {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	res := allPathsSourceTarget(graph)
	fmt.Println(res)
}
