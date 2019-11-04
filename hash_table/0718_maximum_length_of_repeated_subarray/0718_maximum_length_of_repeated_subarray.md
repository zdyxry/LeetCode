718. Maximum Length of Repeated Subarray


Medium


Given two integer arrays A and B, return the maximum length of an subarray that appears in both arrays.

Example 1:

```
Input:
A: [1,2,3,2,1]
B: [3,2,1,4,7]
Output: 3
Explanation: 
The repeated subarray with maximum length is [3, 2, 1].
```

Note:

1 <= len(A), len(B) <= 1000

0 <= A[i], B[i] < 100


## 方法


```go
func findLength(a1 []int, a2 []int) int {
    n1, n2 := len(a1), len(a2)

	// dp[i][j] == k  表示 a1[i-k:i] == a2[j-k:j] ，
	// 但是 a1[i-k-1] != a2[j-k-1]
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}

	res := 0
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if a1[i-1] == a2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				res = max(res, dp[i][j])
			}
		}
	}

	return res
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
    def findLength(self, A, B):
        """
        :type A: List[int]
        :type B: List[int]
        :rtype: int
        """
        if len(A) < len(B): return self.findLength(B, A)
        result = 0
        dp = [[0] * (len(B)+1) for _ in xrange(2)]
        for i in xrange(len(A)):
            for j in xrange(len(B)):
                if A[i] == B[j]:
                    dp[(i+1)%2][j+1] = dp[i%2][j]+1
                else:
                    dp[(i+1)%2][j+1] = 0
            result = max(result, max(dp[(i+1)%2]))
        return result
```