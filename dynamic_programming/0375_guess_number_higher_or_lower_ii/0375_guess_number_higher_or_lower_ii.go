package main

import (
	"fmt"
)

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for j := 2; j <= n; j++ {
		for i := j - 1; 0 < i; i-- {
			dp[i][j] = i + dp[i+1][j]
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], k+max(dp[i][k-1], dp[k+1][j]))
			}
		}
	}

	return dp[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	n := 1
	res := getMoneyAmount(n)
	fmt.Println(res)
}
