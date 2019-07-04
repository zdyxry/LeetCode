415. Add Strings

Easy


Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:

The length of both num1 and num2 is < 5100.

Both num1 and num2 contains only digits 0-9.

Both num1 and num2 does not contain any leading zero.

You must not use any built-in BigInteger library or convert the inputs to integer directly.


## 方法

计算两个字符串长度，从最后位开始相加，考虑进位情况；最后计算完成后，考虑结果反转情况。



```go
func addStrings(num1 string, num2 string) string {
    if len(num1) > len(num2) {
		num1, num2 = num2, num1
	}

	n1, n2 := len(num1), len(num2)
	a1, a2 := []byte(num1), []byte(num2)

	carry := byte(0)

	buf := make([]byte, n2+1)
	buf[0] = '1'

	i := 1
	for i <= n2 {
		if i <= n1 {
			buf[n2+1-i] = a1[n1-i] - '0'
		}
		buf[n2+1-i] += a2[n2-i] + carry

		if buf[n2+1-i] > '9' {
			buf[n2+1-i] -= 10
			carry = byte(1)
		} else {
			carry = byte(0)
		}

		i++
	}

	if carry == 1 {
		return string(buf)
	}
	return string(buf[1:])

}
```




```python
class Solution(object):
    def addStrings(self, num1, num2):
        """
        :type num1: str
        :type num2: str
        :rtype: str
        """
        result = []
        i, j, carry = len(num1) - 1, len(num2) - 1, 0

        while i >= 0 or j >= 0 or carry:
            if i >= 0:
                carry += ord(num1[i]) - ord('0')
                i -= 1
            if j >= 0:
                carry += ord(num2[j]) - ord('0')
                j -= 1
            result.append(str(carry % 10))
            carry /= 10
        result.reverse()

        return "".join(result)
```
