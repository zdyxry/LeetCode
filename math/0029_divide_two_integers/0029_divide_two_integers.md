29. Divide Two Integers


Medium


Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:

```
Input: dividend = 10, divisor = 3
Output: 3
```

Example 2:

```
Input: dividend = 7, divisor = -3
Output: -2
```

Note:

Both dividend and divisor will be 32-bit signed integers.  
The divisor will never be 0.  
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.   


## 方法

```go
func divide(m, n int) int {
	// 防止有人把0当做除数
	if n == 0 {
		return math.MaxInt32
	}

	signM, absM := analysis(m)
	signN, absN := analysis(n)

	res, _ := d(absM, absN, 1)

	// 修改res的符号
	if signM != signN {
		res = res - res - res
	}

	// 检查溢出
	if res < math.MinInt32 || res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}

func analysis(num int) (sign, abs int) {
	sign = 1
	abs = num
	if num < 0 {
		sign = -1
		abs = num - num - num
	}

	return
}

// d 计算m/n的值，返回结果和余数
// m >= 0
// n > 0
// count == 1, 代表初始n的个数，在递归过程中，count == 当前的n/初始的n
func d(m, n, count int) (res, remainder int) {
	switch {
	case m < n:
		return 0, m
	case n <= m && m < n+n:
		return count, m - n
	default:
		res, remainder = d(m, n+n, count+count)
		if remainder >= n {
			return res + count, remainder - n
		}
		return res, remainder
	}
}
```


```python
class Solution(object):
    def divide(self, dividend, divisor):
        """
        :type dividend: int
        :type divisor: int
        :rtype: int
        """
        positive = (dividend < 0) is (divisor < 0)
        dividend, divisor = abs(dividend), abs(divisor)
        res = 0
        while dividend >= divisor:
            temp, i = divisor, 1
            while dividend >= temp:
                dividend -= temp
                res += i
                i <<= 1
                temp <<= 1
        if not positive:
            res = -res
        return min(max(-2147483648, res), 2147483647)
```