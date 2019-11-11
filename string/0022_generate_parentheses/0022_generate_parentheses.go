package main 

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	res := make([]string, 0, n*n)
	bytes := make([]byte, n * 2)
	dfs(n, n, 0, bytes, &res)
	return res
}

func dfs(left, right, idx int, bytes []byte, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, string(bytes))
		return
	}
	if left > 0 {
		bytes[idx] = '('
		dfs(left-1, right, idx+1, bytes, res)
	}

	if right > 0 && left < right {
		bytes[idx] = ')'
		dfs(left, right-1, idx+1, bytes, res)
	}
}

func main() {
	n := 3
	res := generateParenthesis(n)
	fmt.Println(res)
}