171. Excel Sheet Column Number


Easy


Given a column title as appear in an Excel sheet, return its corresponding column number.

For example:

```
    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28 
    ...
```


Example 1:

```
Input: "A"
Output: 1
```

Example 2:
```
Input: "AB"
Output: 28
```

Example 3:
```
Input: "ZY"
Output: 701
```


## 方法


```go
func titleToNumber(s string) int {
    res := 0

	for i := 0; i < len(s); i++ {
		temp := int(s[i] - 'A' + 1)
		res = res*26 + temp
	}

	return res
}
```


```python
class Solution(object):
    def titleToNumber(self, s):
        """
        :type s: str
        :rtype: int
        """
        result = 0
        for i in xrange(len(s)):
            result *= 26
            result += ord(s[i]) - ord('A') + 1
        return result
```