1638. Count Substrings That Differ by One Character


Medium


Given two strings s and t, find the number of ways you can choose a non-empty substring of s and replace a single character by a different character such that the resulting substring is a substring of t. In other words, find the number of substrings in s that differ from some substring in t by exactly one character.

For example, the underlined substrings in "computer" and "computation" only differ by the 'e'/'a', so this is a valid way.

Return the number of substrings that satisfy the condition above.

A substring is a contiguous sequence of characters within a string.

 

Example 1:

```
Input: s = "aba", t = "baba"
Output: 6
Explanation: The following are the pairs of substrings from s and t that differ by exactly 1 character:
("aba", "baba")
("aba", "baba")
("aba", "baba")
("aba", "baba")
("aba", "baba")
("aba", "baba")
The underlined portions are the substrings that are chosen from s and t.
```

​​Example 2:
```
Input: s = "ab", t = "bb"
Output: 3
Explanation: The following are the pairs of substrings from s and t that differ by 1 character:
("ab", "bb")
("ab", "bb")
("ab", "bb")
​​​​The underlined portions are the substrings that are chosen from s and t.
```

Example 3:
```
Input: s = "a", t = "a"
Output: 0
```

Example 4:

```
Input: s = "abe", t = "bbc"
Output: 10
```
 

Constraints:

1 <= s.length, t.length <= 100   
s and t consist of lowercase English letters only.


## 方法


```go
func countSubstrings(s string, t string) int {
	n := len(s)
	m := len(t)
	res := 0
	var k, l int
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			diff := 0
			k = i
			l = j
			for k < n && l < m && diff <= 1 {
				if s[k] != t[l] {
					diff++
				}
				if diff == 1 {
					res++
				}
				k++
				l++
			}
		}
	}
	return res
}
```


```python
class Solution:
    def countSubstrings(self, s: str, t: str) -> int:
        n = len(s)
        m = len(t)
        res = 0
        for i in range(n):
            for j in range(m):
                diff = 0
                k = i
                l = j
                while k < n and l < m and diff <= 1:
                    if s[k] != t[l]:
                        diff += 1
                    
                    if diff == 1:
                        res += 1
                    
                    k += 1
                    l += 1
        
        return res
```


```python
class Solution:
    def countSubstrings(self, s: str, t: str) -> int:
        # 相同长度下s1s2是否满足只差一个字符
        def isvalid(s1,s2):
            diff = 0
            for i in range(len(s1)):
                if s1[i] != s2[i]:
                    diff += 1
            return True if diff == 1 else False

        # 预处理s所有子串
        substr = []
        for i in range(len(s)):
            for j in range(len(s) - i):
                substr.append(s[i:i + j + 1])
        ans = 0
        n = len(t)
        for w in substr:
            l = len(w)
            for left in range(len(t) - l + 1):
                if isvalid(w,t[left:left + l]):
                    ans += 1
        return ans
```