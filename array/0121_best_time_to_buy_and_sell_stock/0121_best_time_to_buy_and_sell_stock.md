121. Best Time to Buy and Sell Stock

Easy

Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:
```
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
```

Example 2:
```
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
```

# 方法
1. 遍历取当前最小值，最大利益为当前值减去最小值与当前最大利益值取大

2. https://www.polarxiong.com/archives/LeetCode-121-best-time-to-buy-and-sell-stock.html



```go
func maxProfit(prices []int) int {
	max := 0
	temp := 0
	for i := 1; i < len(prices); i++ {
		temp += prices[i] - prices[i-1]
		if temp < 0 {
			temp = 0
		}
		if max < temp {
			max = temp
		}
	}

	return max
}
```

```go
func maxProfit(prices []int) int {
    
    if len(prices) == 0{
        return 0
    }
    min_price := prices[0]
    max_profit := 0
    for _, price := range(prices) {
        min_price = min(price, min_price)
        max_profit = max(max_profit, price - min_price)
	}

	return max_profit
}

func max(a int, b int) int {
    if a > b{
        return a
    } else {
        return b
    }
}

func min(a int, b int) int {
    if a < b{
        return a
    } else {
        return b
    }
}
```




```python
class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        min_price, max_profit = float('inf'), 0
        for price in prices:
            min_price = min(price, min_price)
            max_profit = max(max_profit, price - min_price)
        return max_profit
```