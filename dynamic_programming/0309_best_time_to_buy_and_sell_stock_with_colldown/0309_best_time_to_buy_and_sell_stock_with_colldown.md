309. Best Time to Buy and Sell Stock with Cooldown


Medium

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:

    You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
    After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)

Example:

```
Input: [1,2,3,0,2]
Output: 3 
Explanation: transactions = [buy, sell, cooldown, buy, sell]
```

## 方法 

```go
func maxProfit(prices []int) int {
    n := len(prices)
	if n == 0 {
		return 0
	}

	// prices[i] 表示 第 i 天的价格
	buy := make([]int, n+1)
	buy[1] = 0 - prices[0]
	sel := make([]int, n+1)

	for i := 2; i <= n; i++ {
		buy[i] = max(buy[i-1], sel[i-2]-prices[i-1])
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
```


```python
class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        n = len(prices)
        if n < 2:
            return 0
        if n == 2:
            return max(0, prices[n-1] - prices[n-2])
        dp_prev_prev = 0
        dp_prev = max(0, prices[n-1] - prices[n-2])
        dp_curr = max(dp_prev, prices[n-2] - prices[n-3], prices[n-1] - prices[n-3])
        max_so_far = max(prices[n-1], prices[n-2])
        for i in xrange(n-4,-1,-1):
            max_so_far = max(max_so_far, prices[i+1] + dp_prev_prev)
            dp_prev_prev = dp_prev
            dp_prev = dp_curr
            dp_curr = max(dp_curr, max_so_far - prices[i])
        return dp_curr
```