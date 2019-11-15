647. Palindromic Substrings


Medium


Given a string, your task is to count how many palindromic substrings in this string.

The substrings with different start indexes or end indexes are counted as different substrings even they consist of same characters.

Example 1:

```
Input: "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
```
 

Example 2:

```
Input: "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
```

Note:

The input string length won't exceed 1000.


## 方法

```go
func countSubstrings(s string) int {
    count := 0
	// 换一个思路
	for i := 0; i < len(s); i++ {
		// 统计所有以 i 为中心的回文
		count += extendPalindrome(s, i, i)
		// 统计所有以 i 和 i+1 为中心的回文
		count += extendPalindrome(s, i, i+1)
	}
	return count
}

// 同时向两边扩张，直到遇到不是回文的情况
func extendPalindrome(s string, left, right int) int {
	res := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		res++
		left--
		right++
	}
	return res
}
```



```python
class Solution(object):
    def countSubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        def manacher(s):
            s = '^#' + '#'.join(s) + '#$'
            P = [0] * len(s)
            C, R = 0, 0
            for i in xrange(1, len(s) - 1):
                i_mirror = 2*C-i
                if R > i:
                    P[i] = min(R-i, P[i_mirror])
                while s[i+1+P[i]] == s[i-1-P[i]]:
                    P[i] += 1
                if i+P[i] > R:
                    C, R = i, i+P[i]
            return P
        return sum((max_len+1)//2 for max_len in manacher(s))

```