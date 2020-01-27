package main

import (
	"fmt"
)

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	dp := make([][]int, 2)
	dp[0] = make([]int, len(matrix[0]))
	dp[1] = make([]int, len(matrix[0]))
	maxLen := 0
	for i := range matrix[0] {
		dp[0][i] = int(matrix[0][i] - '0')
		dp[1][i] = int(matrix[0][i] - '0')
		maxLen = maxLen + dp[1][i] - min(maxLen, dp[1][i])

	}

	fmt.Println(maxLen)
	for i := 1; i < len(matrix); i++ {
		dp[1][0] = int(matrix[i][0] - '0')
		maxLen = maxLen + dp[1][0] - min(maxLen, dp[1][0])
		for j := 1; j < len(matrix[0]); j++ {
			dp[1][j] = 0
			if int(matrix[i][j]-'0') == 1 {
				dp[1][j] = 1 + min(dp[0][j-1], min(dp[0][j], dp[1][j-1]))
			}
			maxLen = maxLen + dp[1][j] - min(maxLen, dp[1][j])
		}
		for i, val := range dp[1] {
			dp[0][i] = val
		}
	}
	return maxLen * maxLen
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	res := maximalSquare(matrix)
	fmt.Println(res)
}
