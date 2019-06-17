409. Longest Palindrome

Easy

Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:

Assume the length of given string will not exceed 1,010.

Example:

```
Input:
"abccccdd"

Output:
7

Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.
```


# 方法

遍历字符串，记录每个字符串出现的次数，若出现为偶数次，则直接相加，若为奇数次，则加次数-1，最终加上 1 个奇数。

关于是否实用 map 来实现，观看 ： https://www.youtube.com/watch?v=qdpBD0LFgN0 


```go
func longestPalindrome(s string) int {
    res := 0
	a := [123]int{} // 'z' 的 ASCII 码为 122
	for i := range s {
		a[s[i]]++
	}

	hasOdd := 0
	for i := range a {
		if a[i] == 0 {
			continue
		}
		if a[i]&1 == 0 {
			res += a[i]
		} else {
			res += a[i] - 1
			hasOdd = 1
		}
	}

	return res + hasOdd
}
```


```python
class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: int
        """
        if s is None or len(s) == 0:
            return 0
        dic = dict()
        for c in s:
            if dic.get(c) is None:
                dic[c] = 1
            else:
                dic[c] += 1
        f = 0
        ans = 0
        for key, val in dic.items():
            if val % 2 == 0:
                ans += val
            elif val == 1:
                f = 1
            else:
                ans += val-1
                f = 1
        return ans+f

```


```python
class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: int
        """
        odds = sum(v & 1 for v in collections.Counter(s).values())
        return len(s) - odds + bool(odds)
```