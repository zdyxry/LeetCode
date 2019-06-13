242. Valid Anagram

Easy

Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:
```
Input: s = "anagram", t = "nagaram"
Output: true
```
Example 2:

```
Input: s = "rat", t = "car"
Output: false
```

Note:

You may assume the string contains only lowercase alphabets.

Follow up:

What if the inputs contain unicode characters? How would you adapt your solution to such case?


# 方法
遍历两个字符串，记录出现过的字母及次数，比较是否相等。



```go
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}

	m := make([]int, 26)
	for i, x := range s {
		m[x-'a'] ++
		m[t[i]-'a'] --
	}

	for _, i := range m {
		if i != 0 {
			return false
		}
	}
	return true
}
```


```python
class Solution(object):
    def isAnagram(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        dic1, dic2 = {}, {}
        for item in s:
            dic1[item] = dic1.get(item, 0) + 1
        for item in t:
            dic2[item] = dic2.get(item, 0) + 1
        return dic1 == dic2
```


```python
class Solution(object):
    def isAnagram(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        dic1, dic2 = [0]*26, [0]*26
        for item in s:
            dic1[ord(item)-ord('a')] += 1
        for item in t:
            dic2[ord(item)-ord('a')] += 1
        return dic1 == dic2

```