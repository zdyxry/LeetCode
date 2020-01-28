package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	buy := make([]int, n+1)
	buy[1] = 0 - prices[0]
	sel := make([]int, n+1)

	for i := 2; i <= n; i++ {
		// buy[i] 最佳收益 = max(昨天买入的最佳收益，前天卖出的最佳收益 - 第 i 天价格)
		buy[i] = max(buy[i-1], sel[i-2]-prices[i-1])
		// sel[i] 最佳收益 = max(昨天卖出的最佳收益，昨天买入的最佳收益 + 第 i 天价格)
		sel[i] = max(sel[i-1], buy[i-1]+prices[i-1])
	}
	return sel[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	res := maxProfit(prices)
	fmt.Println(res)
}
