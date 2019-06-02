746. Min Cost Climbing Stairs

Easy

On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

Example 1:
```j
Input: cost = [10, 15, 20]
Output: 15
Explanation: Cheapest is start on cost[1], pay that cost and go to the top.
```

Example 2:
```
Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
Output: 6
Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].
```

Note:

cost will have a length in the range [2, 1000].

Every cost[i] will be an integer in the range [0, 999].


# 方法
1. 动态规划，dp[i] 表示，爬到第 i 个台阶需要的花费，从第 2 个台阶开始遍历，当前台阶的花费等于到达 dp[i-1] 的钱加上 i-1 自身需要的花费 与 达到 dp[i-2] 的钱加上 i-2 自身需要的钱的最小值。
2. 动态规划，dp[i] 表示，爬到第 i 个台阶且离开了第 i 个台阶需要的花费。

```go
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
	// dp[i] 表示，爬到第 i 个台阶需要的花费
	dp := make([]int, n+1)

	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```


```python
class Solution(object):
    def minCostClimbingStairs(self, cost):
        """
        :type cost: List[int]
        :rtype: int
        """
        dp = [0] * 3
        for i in reversed(xrange(len(cost))):
            dp[i%3] = cost[i] + min(dp[(i+1)%3], dp[(i+2)%3])
        return min(dp[0], dp[1])
        
```