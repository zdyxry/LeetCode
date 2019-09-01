367. Valid Perfect Square


Easy


Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:

```
Input: 16
Output: true
```

Example 2:

```
Input: 14
Output: false
```


## 方法

```go
func isPerfectSquare(num int) bool {
    r := intSqrt(num)
	return r*r == num
}

func intSqrt(n int) int {
	r := n

	for r*r > n {
		r = (r + n/r) / 2
	}

	return r
}

```


```python
class Solution(object):
    def isPerfectSquare(self, num):
        """
        :type num: int
        :rtype: bool
        """
        left, right = 1, num
        while left <= right:
            mid = left + (right - left) / 2
            if mid >= num / mid:
                right = mid - 1
            else:
                left = mid + 1
        return left == num / left and num % left == 0
```