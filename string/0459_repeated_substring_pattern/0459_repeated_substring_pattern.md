459. Repeated Substring Pattern

Easy


Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies of the substring together. You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.

 

Example 1:
```
Input: "abab"
Output: True
Explanation: It's the substring "ab" twice.
```

Example 2:
```
Input: "aba"
Output: False
```

Example 3:
```
Input: "abcabcabcabc"
Output: True
Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)
```


## 方法

1. 因为是找重复子字符串，那么最少也有2个重复字符串，所以只需要遍历一半长度即可，当索引可以被整除时，将子字符串乘以对应次数，判断是否与原字符串相等。

2. https://leetcode.com/problems/repeated-substring-pattern/discuss/94334/easy-python-solution-with-explaination
 
看了好一会，没看明白，把结果记录下来，也许之后再看会有所体会吧。




```go
func repeatedSubstringPattern(s string) bool {
    sz := len(s)
	if sz <= 1 {
		return false
	}

	for i := 2; i <= len(s); i++ {
		if sz%i == 0 {
			n := sz / i
			match := true
			str1 := s[:n]
			for j := n; j < len(s); j += n {
				if s[j:j+n] != str1 {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}
```

```go
func repeatedSubstringPattern(s string) bool {
    if len(s) == 0 {
		return false
	}

	size := len(s)

	ss := (s + s)[1 : size*2-1]

	return strings.Contains(ss, s)
}
```




```python
class Solution(object):
    def repeatedSubstringPattern(self, s):
        """
        :type s: str
        :rtype: bool
        """
        for l in range(1,len(s)//2+1):
            if len(s)%l ==0:
                str_sub = s[:l]
                n = len(s)//l
                if str_sub*n == s:
                    return True
        return False

```


```python
class Solution(object):
    def repeatedSubstringPattern(self, s):
        """
        :type s: str
        :rtype: bool
        """
        if not str:
            return False

        ss = (s + s)[1:-1]
        print(ss)
        return ss.find(s) != -1
```