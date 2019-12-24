343. Integer Break


Medium


Given a positive integer n, break it into the sum of at least two positive integers and maximize the product of those integers. Return the maximum product you can get.

Example 1:

```
Input: 2
Output: 1
Explanation: 2 = 1 + 1, 1 × 1 = 1.
```

Example 2:

```
Input: 10
Output: 36
Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
```

Note: 

You may assume that n is not less than 2 and not larger than 58.


## 方法

```go
func integerBreak(n int) int {
    dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			// dp[i] = max(dp[i], j * (i - j), j*dp[i-j])
			dp[i] = max(dp[i], j*max(dp[i-j], i-j))
		}
	}
	return dp[n]
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
} 
```


```go
func integerBreak(n int) int {
	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	switch n % 3 {
	case 0:
		return pow3(n / 3)
	case 1:
		return 4 * pow3(n/3-1)
	default:
		return 2 * pow3(n/3)
	}

}

func pow3(n int) int {
	if n == 0 {
		return 1
	}

	res := pow3(n >> 1)
	if n&1 == 0 {
		return res * res
	}

	return res * res * 3
}
```





```python
class Solution(object):
    def integerBreak(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n < 4:
            return n - 1

        #  Proof.
        #  1. Let n = a1 + a2 + ... + ak, product = a1 * a2 * ... * ak
        #      - For each ai >= 4, we can always maximize the product by:
        #        ai <= 2 * (ai - 2)
        #      - For each aj >= 5, we can always maximize the product by:
        #        aj <= 3 * (aj - 3)
        #
        #     Conclusion 1:
        #      - For n >= 4, the max of the product must be in the form of
        #        3^a * 2^b, s.t. 3a + 2b = n
        #
        #  2. To maximize the product = 3^a * 2^b s.t. 3a + 2b = n
        #      - For each b >= 3, we can always maximize the product by:
        #        3^a * 2^b <= 3^(a+2) * 2^(b-3) s.t. 3(a+2) + 2(b-3) = n
        #
        #     Conclusion 2:
        #      - For n >= 4, the max of the product must be in the form of
        #        3^Q * 2^R, 0 <= R < 3 s.t. 3Q + 2R = n
        #        i.e.
        #          if n = 3Q + 0,   the max of the product = 3^Q * 2^0
        #          if n = 3Q + 2,   the max of the product = 3^Q * 2^1
        #          if n = 3Q + 2*2, the max of the product = 3^Q * 2^2

        res = 0
        if n % 3 == 0:            #  n = 3Q + 0, the max is 3^Q * 2^0
            res = 3 ** (n // 3)
        elif n % 3 == 2:          #  n = 3Q + 2, the max is 3^Q * 2^1
            res = 3 ** (n // 3) * 2
        else:                     #  n = 3Q + 4, the max is 3^Q * 2^2
            res = 3 ** (n // 3 - 1) * 4
        return res
```