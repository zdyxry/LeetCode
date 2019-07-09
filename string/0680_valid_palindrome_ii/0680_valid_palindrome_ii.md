680. Valid Palindrome II



Easy


Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

Example 1:
```
Input: "aba"
Output: True
```

Example 2:
```
Input: "abca"
Output: True
Explanation: You could delete the character 'c'.
```

Note:

The string will only contain lowercase characters a-z. The maximum length of the string is 50000.

## 方法

双指针，当两个指针所指的字符不相等时，删掉低位指针或高位指针所指向的字符，继续判断是否为回文。

递归。




```go
func validPalindrome(s string) bool {
	return helper([]byte(s), 0, len(s)-1, false)
}

func helper(bs []byte, lo, hi int, hasDeleted bool) bool {
	for lo < hi {
		if bs[lo] != bs[hi] {
			if hasDeleted {
				return false
			}
			return helper(bs, lo+1, hi, true) || // 删除 s[lo]
				helper(bs, lo, hi-1, true) // 删除 s[hi]
		}
		lo++
		hi--
	}

	return true
}
```





```python
class Solution(object):
    def validPalindrome(self, s):
        """
        :type s: str
        :rtype: bool
        """
        def validPalindrome(s, left, right):
            while left < right:
                if s[left] != s[right]:
                    return False
                left, right = left+1, right-1
            return True

        left, right = 0, len(s)-1
        while left < right:
            if s[left] != s[right]:
                return validPalindrome(s, left, right-1) or validPalindrome(s, left+1, right)
            left, right = left+1, right-1
        return True

```