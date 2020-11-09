package main

import "fmt"

func countVowelStrings(n int) int {
	dp := make([][5]int, n)

	for i := 0; i < 5; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k <= j; k++ {
				dp[i][j] += dp[i-1][k]
			}
		}
	}

	count := 0

	for i := 0; i < 5; i++ {
		count += dp[n-1][i]
	}

	return count
}

func main() {
	n := 1
	res := countVowelStrings(n)
	fmt.Println(res)
}
