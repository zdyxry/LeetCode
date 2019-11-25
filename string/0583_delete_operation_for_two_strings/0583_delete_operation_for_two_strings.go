package main 

import (
	"fmt"
)

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1) + 1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2) + 1)
	}

	for i := 0; i <= len(word1); i++ {
		for j := 0; j<= len(word2); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = i + j 
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func Min(x, y int) int {
	if x < y {
		return x 
	}
	return y
}

func main() {
	word1, word2 := "sea", "eat"
	res := minDistance(word1, word2)
	fmt.Println(res)
}