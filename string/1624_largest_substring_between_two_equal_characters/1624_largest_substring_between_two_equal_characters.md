1624. Largest Substring Between Two Equal Characters


Easy


Given a string s, return the length of the longest substring between two equal characters, excluding the two characters. If there is no such substring return -1.

A substring is a contiguous sequence of characters within a string.

 

Example 1:

```
Input: s = "aa"
Output: 0
Explanation: The optimal substring here is an empty substring between the two 'a's.
```

Example 2:

```
Input: s = "abca"
Output: 2
Explanation: The optimal substring here is "bc".
```

Example 3:

```
Input: s = "cbzxy"
Output: -1
Explanation: There are no characters that appear twice in s.
```

Example 4:

```
Input: s = "cabbac"
Output: 4
Explanation: The optimal substring here is "abba". Other non-optimal substrings include "bb" and "".
```

Constraints:

1 <= s.length <= 300   
s contains only lowercase English letters.


## 方法


```go
func maxLengthBetweenEqualCharacters(s string) int {
    m := make(map[rune]int)
	max := -1
	for i, ch := range s {
		if pos, ok := m[ch]; ok {
			if i-pos-1 > max {
				max = i - pos - 1
			}
		} else {
			m[ch] = i
		}
	}
	return max
}
```


```python
class Solution:
    def maxLengthBetweenEqualCharacters(self, s: str) -> int:
        left = {}
        right = {}
        res = -1
        for i, c in enumerate(s):
            if c not in left:
                left[c] = i
            else:
                right[c] = i
                res = max(res, right[c] - left[c]-1)
        return res
```