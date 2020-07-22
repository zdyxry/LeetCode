package main

import "fmt"

func countSubTrees(n int, edges [][]int, labels string) []int {
	relations := genRelations(edges)
	ans := make([]int, n)

	dfs(0, relations, labels, -1, &ans)
	return ans
}

func dfs(i int, relations [][]int, labels string, visited int, ans *[]int) []int {
	count := make([]int, 26)
	count[labels[i]-'a']++

	for _, c := range relations[i] {
		if c == visited {
			continue
		}

		re := dfs(c, relations, labels, i, ans)
		for i := 0; i < 26; i++ {
			count[i] += re[i]
		}
	}

	(*ans)[i] = count[labels[i]-'a']
	return count
}

func genRelations(edges [][]int) [][]int {
	relations := make([][]int, len(edges)+1)
	for _, e := range edges {
		relations[e[0]] = append(relations[e[0]], e[1])
		relations[e[1]] = append(relations[e[1]], e[0])
	}

	return relations
}

func main() {
	n := 7
	edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	labels := "abaedcd"
	res := countSubTrees(n, edges, labels)
	fmt.Println(res)
}
