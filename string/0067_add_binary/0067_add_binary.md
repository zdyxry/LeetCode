67. Add Binary


Easy


Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:
```
Input: a = "11", b = "1"
Output: "100"
```

Example 2:
```
Input: a = "1010", b = "1011"
Output: "10101"
```

## 方法

从最低位开始相加，考虑进位情况。

python 可以直接通过 int(a, 2)转换为二进制相加解决。



```go
func addBinary(a string, b string) string {
	carry, ans := 0, ""
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		if sum >= 2 {
			carry = 1
			sum -= 2
		} else {
			carry = 0
		}
		ans = strconv.Itoa(sum) + ans
	}
	return ans
}
```


```python
class Solution(object):
    def addBinary(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        result, carry, val = "", 0, 0
        for i in xrange(max(len(a), len(b))):
            val = carry
            if i < len(a):
                val += int(a[-(i + 1)])
            if i < len(b):
                val += int(b[-(i + 1)])
            carry, val = divmod(val, 2)
            result += str(val)
        if carry:
            result += str(carry)
        return result[::-1]
```