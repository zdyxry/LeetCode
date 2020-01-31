package main

import (
	"fmt"
	"math"
)

func numSquares1(n int) int {
	var ans int
	seen := make([]bool, n)
	q := []int{n}
	for len(q) > 0 {
		ans++
		var newQ []int
		for _, x := range q {
			for i := 1; i*i <= x; i++ {
				if x == i*i {
					return ans
				}
				if !seen[x-i*i] {
					newQ = append(newQ, x-i*i)
					seen[x-i*i] = true
				}
			}
		}
		q = newQ
	}
	return ans
}

func numSquares2(n int) int {
	perfects := []int{}
	for i := 1; i*i <= n; i++ {
		perfects = append(perfects, i*i)
	}

	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for _, p := range perfects {
		for i := p; i < len(dp); i++ {
			if dp[i] > dp[i-p]+1 {
				dp[i] = dp[i-p] + 1
			}
		}
	}

	return dp[n]
}

func main() {
	n := 12
	res := numSquares1(n)
	fmt.Println(res)
}
