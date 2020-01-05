50. Pow(x, n)


Medium


Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

```
Input: 2.00000, 10
Output: 1024.00000
```

Example 2:

```
Input: 2.10000, 3
Output: 9.26100
```

Example 3:

```
Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
```

Note:

-100.0 < x < 100.0   
n is a 32-bit signed integer, within the range [−231, 231 − 1]

## 方法

```go
func myPow(x float64, n int) float64 {
    if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}
```

```go
func myPow(x float64, n int) float64 {
    if n < 0 {
		return 1.0 / pow(x, -n)
	}

	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	res := pow(x, n>>1)
	if n&1 == 0 {
		return res * res
	}

	return res * res * x
}
```



```python
class Solution(object):
    def myPow(self, x, n):
        """
        :type x: float
        :type n: int
        :rtype: float
        """
        result = 1
        abs_n = abs(n)
        while abs_n:
            if abs_n & 1:
                result *= x
            abs_n >>= 1
            x *= x

        return 1 / result if n < 0 else result
```