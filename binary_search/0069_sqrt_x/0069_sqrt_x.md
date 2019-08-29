69. Sqrt(x)


Easy


Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

```
Input: 4
Output: 2
```


Example 2:
```
Input: 8
Output: 2

Explanation: The square root of 8 is 2.82842..., and since 
             the decimal part is truncated, 2 is returned.
```

## 方法

```go
func mySqrt(x int) int {
    if x == 0 {
		return 0
	}
	left, right, res := 1, x, 0
	for left <= right {
		mid := left + ((right - left) >> 1)
		if mid < x/mid {
			left = mid + 1
			res = mid
		} else if mid == x/mid {
			return mid
		} else {
			right = mid - 1
		}
	}
	return res
}
```



```python
class Solution(object):
    def mySqrt(self, x):
        """
        :type x: int
        :rtype: int
        """
        if x < 2:
            return x

        left, right = 1, x // 2
        while left <= right:
            mid = left + (right - left) // 2
            if mid > x / mid:
                right = mid - 1
            else:
                left = mid + 1

        return left - 1
```