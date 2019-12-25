172. Factorial Trailing Zeroes


Easy


Given an integer n, return the number of trailing zeroes in n!.

Example 1:

```
Input: 3
Output: 0
Explanation: 3! = 6, no trailing zero.
```

Example 2:

```
Input: 5
Output: 1
Explanation: 5! = 120, one trailing zero.
```

Note: 

Your solution should be in logarithmic time complexity.

## 方法

```go
func trailingZeroes(n int) int {
    res := 0

	for n >= 5 {
		n /= 5
		res += n
	}

	return res
}
```


```python
class Solution(object):
    def trailingZeroes(self, n):
        """
        :type n: int
        :rtype: int
        """
        if n < 5: return 0

        return n//5 + self.trailingZeroes(n//5)
```
