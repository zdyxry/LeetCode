1763. Longest Nice Substring


Easy


A string s is nice if, for every letter of the alphabet that s contains, it appears both in uppercase and lowercase. For example, "abABB" is nice because 'A' and 'a' appear, and 'B' and 'b' appear. However, "abA" is not because 'b' appears, but 'B' does not.

Given a string s, return the longest substring of s that is nice. If there are multiple, return the substring of the earliest occurrence. If there are none, return an empty string.

 

Example 1:

```
Input: s = "YazaAay"
Output: "aAa"
Explanation: "aAa" is a nice string because 'A/a' is the only letter of the alphabet in s, and both 'A' and 'a' appear.
"aAa" is the longest nice substring.
```

Example 2:

```
Input: s = "Bb"
Output: "Bb"
Explanation: "Bb" is a nice string because both 'B' and 'b' appear. The whole string is a substring.
```

Example 3:

```
Input: s = "c"
Output: ""
Explanation: There are no nice substrings.
```

Example 4:

```
Input: s = "dDzeE"
Output: "dD"
Explanation: Both "dD" and "eE" are the longest nice substrings.
As there are multiple longest nice substrings, return "dD" since it occurs earlier.
```

Constraints:

1 <= s.length <= 100    
s consists of uppercase and lowercase English letters.


## 方法

```go
func longestNiceSubstring(s string) string {
	check := func(s string) bool {
		counts := map[rune]int{}
		for _, ch := range s {
			if ch <= 'Z' {
				ch += 'a' - 'A'
				counts[ch] |= 1
			} else {
				counts[ch] |= 2
			}
		}

		for _, count := range counts {
			if count != 3 {
				return false
			}
		}

		return true
	}

	for i := 0; i < len(s)-1; i++ {
		for j := 0; j <= i; j++ {
			sub := s[j : len(s)-i+j]

			if check(sub) {
				return sub
			}
		}
	}
	return ""
}

```




```python
class Solution:
    def longestNiceSubstring(self, s: str) -> str:
        res = ''
        n = len(s)
        for i in range(n):
            for j in range(i + 1, n):
                t = s[i:j + 1]
                if all(c.upper() in t and c.lower() in t for c in t):
                    if len(t) > len(res):
                        res = t
        return res
```