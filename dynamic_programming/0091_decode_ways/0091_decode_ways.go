package main

import (
	"fmt"
	"strconv"
)

func numDecodings1(s string) int {
	if s[0] == '0' {
		return 0
	}

	n := len(s)
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	lastOne, lastTwo := int(s[0]-'0'), 0
	for i := 2; i <= n; i++ {
		last := int(s[i-1] - '0')
		lastOne, lastTwo = last, lastOne*10+last
		if 1 <= lastOne {
			dp[i] = dp[i-1]
		}
		if 10 <= lastTwo && lastTwo <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

func numDecodings2(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[:1] == "0" {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i <= len(s); i++ {
		lastNum, _ := strconv.Atoi(s[i-1 : i])
		if lastNum >= 1 && lastNum <= 9 {
			dp[i] += dp[i-1]
		}
		lastNum, _ = strconv.Atoi(s[i-2 : i])
		if lastNum >= 10 && lastNum <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

func main() {
	s := "112"
	res := numDecodings2(s)
	fmt.Println(res)
}
