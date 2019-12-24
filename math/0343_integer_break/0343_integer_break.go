package main

import (
	"fmt"
)

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], j*max(dp[i-j], i-j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := 5
	res := integerBreak(n)
	fmt.Println(res)
}
