package main

import "fmt"

func findLexSmallestString(s string, a int, b int) string {

	visited := make(map[string]bool)
	res := s
	add := func(s string) string {
		temp := []byte(s)
		for i := 1; i < len(s); i += 2 {
			temp[i] = byte('0' + (int(temp[i]-'0')+a)%10)
		}
		return string(temp)
	}
	var dfs func(s string)
	dfs = func(s string) {
		if visited[s] {
			return
		}
		visited[s] = true
		if s < res {
			res = s
		}
		dfs(add(s))
		dfs(s[b:] + s[:b])
	}
	dfs(s)

	return res
}
func main() {
	s := "5525"
	a := 9
	b := 2
	res := findLexSmallestString(s, a, b)
	fmt.Println(res)
}
