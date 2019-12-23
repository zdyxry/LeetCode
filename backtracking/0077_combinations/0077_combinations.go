package main 

import (
	"fmt"
)

func combine(n, k int) [][]int {
	combination := make([]int, k)
	res := [][]int{}

	var dfs func(int, int)
	dfs = func(idx, begin int) {
		if idx == k{
			temp := make([]int, k)
			copy(temp, combination)
			res = append(res, temp)
			return
		}
		
		for i := begin; i <= n+1-k+idx; i++ {
			combination[idx] = i
			dfs(idx+1, i+1)
		}
	}

	dfs(0, 1)
	return res
}

func main() {
	n := 4
	k := 2
	res := combine(n, k)
	fmt.Println(res)
}