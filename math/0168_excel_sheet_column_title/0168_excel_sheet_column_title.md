168. Excel Sheet Column Title


Easy


Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:
```
    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB 
    ...

```


Example 1:
```
Input: 1
Output: "A"
```

Example 2:
```
Input: 28
Output: "AB"
```

Example 3:
```
Input: 701
Output: "ZY"

```


## 方法

```go
func convertToTitle(n int) string {
    res := ""

	for n > 0 {
		n--
		res = string(byte(n%26)+'A') + res
		n /= 26
	}

	return res
}
```


```python
class Solution(object):
    def convertToTitle(self, n):
        """
        :type n: int
        :rtype: str
        """
        if n==0:
            return ""
        else:
            r = (n-1)%26
            q = (n-1)//26
            return self.convertToTitle(q) + chr(r + ord('A'))
```