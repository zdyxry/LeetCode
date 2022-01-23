package main

func validPath(n int, edges [][]int, start int, end int) bool {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	visited := make(map[int]bool)
	q := []int{start}
	visited[start] = true
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur == end {
			return true
		}
		for _, node := range graph[cur] {
			if !visited[node] {
				visited[node] = true
				q = append(q, node)
			}
		}
	}
	return false
}
