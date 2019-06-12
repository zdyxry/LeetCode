205. Isomorphic Strings

Easy

Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

Example 1:

```
Input: s = "egg", t = "add"
Output: true
```
Example 2:

```
Input: s = "foo", t = "bar"
Output: false
```
Example 3:

```
Input: s = "paper", t = "title"
Output: true
```

Note:

You may assume both s and t have the same length.

# 方法
遍历字符，分别统计各个字符出现的次数，比较是否一致。


```go
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make([]int, 256)
	m2 := make([]int, 256)

	for i := 0; i < len(s); i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}

		m1[int(s[i])] = i + 1
		m2[int(t[i])] = i + 1
	}
	return true
}
```

```python

class Solution(object):
    def isIsomorphic(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        a = [0] * 256
        b = [0] * 256
        for x in xrange(len(s)):
            if a[ord(s[x])] != b[ord(t[x])]:
                return False
            a[ord(s[x])] = x + 1
            b[ord(t[x])] = x + 1
        return True

```
