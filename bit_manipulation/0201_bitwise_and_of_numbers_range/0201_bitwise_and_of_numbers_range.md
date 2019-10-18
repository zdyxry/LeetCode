201. Bitwise AND of Numbers Range


Medium


Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND of all numbers in this range, inclusive.

Example 1:

```
Input: [5,7]
Output: 4
```

Example 2:

```
Input: [0,1]
Output: 0
```

## 方法


```go
func rangeBitwiseAnd(m int, n int) int {
    res := uint(0)

	for m >= 1 && n >= 1 {
		if m == n {
			return m << res
		}

		m >>= 1
		n >>= 1
		res++
	}

	return 0
}
```


```python
class Solution(object):
    def rangeBitwiseAnd(self, m, n):
        """
        :type m: int
        :type n: int
        :rtype: int
        """
        while m < n:
            n &= n - 1
        return n
```