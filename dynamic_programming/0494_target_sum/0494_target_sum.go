package main

import (
	"fmt"
)

func findTargetSumWays(nums []int, S int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum < S {
		return 0
	}

	if (sum+S)%2 == 1 {
		return 0
	}

	return calcSumWays(nums, (sum+S)/2)

}

func calcSumWays(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}

func main() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	res := findTargetSumWays(nums, target)
	fmt.Println(res)
}
