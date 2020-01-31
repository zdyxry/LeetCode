279. Perfect Squares

Medium

Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

Example 1:

```
Input: n = 12
Output: 3 
Explanation: 12 = 4 + 4 + 4.
```

Example 2:

```
Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
```


## 方法

```go
func numSquares(n int) int {
	perfects := []int{}
	for i := 1; i*i <= n; i++ {
		perfects = append(perfects, i*i)
	}

	// dp[i] 表示 the least number of perfect square numbers which sum to i
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for _, p := range perfects {
		for i := p; i < len(dp); i++ {
			if dp[i] > dp[i-p]+1 {
				// 因为 i = ( i - p ) + p，p 是 平方数
				// 所以 dp[i] = dp[i-p] + 1
				dp[i] = dp[i-p] + 1
			}
		}
	}

	return dp[n]
}
```


```go
func numSquares(n int) int {
    var ans int
    seen := make([]bool, n)
    q := []int{n}
    for len(q) > 0 {
        // 每次遍历均将 ans + 1
        ans++
        var newQ []int
        // 依次遍历q，当 x 无法被开方时，将其加入到 newQ 中，并标记其已计算过
        // 直到 x 可以被开方，则返回 ans
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
```


```python
class Solution(object):
    _num = [0]
    def numSquares(self, n):
        """
        :type n: int
        :rtype: int
        """
        num = self._num
        while len(num) <= n:
            num += min(num[-i*i] for i in xrange(1, int(len(num)**0.5+1))) + 1,
        return num[n]
```

```python
class Solution(object):
    def numSquares(self, n):
        dp=[i for i in range(n+1)]
        for i in range(2,n+1):
            for j in range(1,int(i**(0.5))+1):
                dp[i]=min(dp[i],dp[i-j*j]+1)
        return dp[-1]
```