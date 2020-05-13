package main

func minTime(n int, edges [][]int, hasApple []bool) int {
	graph := make(map[int][]int)
	for _, v := range edges {
		a := v[0]
		b := v[1]
		if _, ok := graph[a]; !ok {
			graph[a] = []int{}
		}
		graph[a] = append(graph[a], b)
		if _, ok := graph[b]; !ok {
			graph[b] = []int{}
		}
		graph[b] = append(graph[b], a)
	}
	set := make(map[int]bool)
	set[0] = true
	return dfs(0, graph, set, hasApple)
}

func dfs(cur int, graph map[int][]int, set map[int]bool, has []bool) int {
	set[cur] = true
	res := 0
	for _, v := range graph[cur] {
		if !set[v] {
			res += dfs(v, graph, set, has)
		}
	}
	if cur == 0 {
		return res
	}
	if res > 0 || has[cur] {
		res += 2
	}
	return res
}
