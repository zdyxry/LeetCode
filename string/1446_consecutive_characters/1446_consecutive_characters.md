1446. Consecutive Characters


Easy


Given a string s, the power of the string is the maximum length of a non-empty substring that contains only one unique character.

Return the power of the string.

 

Example 1:

```
Input: s = "leetcode"
Output: 2
Explanation: The substring "ee" is of length 2 with the character 'e' only.
```

Example 2:

```
Input: s = "abbcccddddeeeeedcba"
Output: 5
Explanation: The substring "eeeee" is of length 5 with the character 'e' only.
```

Example 3:

```
Input: s = "triplepillooooow"
Output: 5
```

Example 4:

```
Input: s = "hooraaaaaaaaaaay"
Output: 11
```

Example 5:

```
Input: s = "tourist"
Output: 1
```
 

Constraints:

1 <= s.length <= 500  
s contains only lowercase English letters.


## 方法

```go
func maxPower(s string) int {
    n, ans := 1, 1
    for i := 1; i < len(s); i++ {
        if s[i - 1] == s[i] {
            n++
        } else {
            n = 1
        }
        if n > ans { ans = n }
    }
    return ans
}
```




```python
class Solution:
    def maxPower(self, s: str) -> int:
        if len(s) == 1:
            return 1
        tmp = s[0]
        cnt = 1
        max_cnt = 0
        for i in s[1:]:
            if i == tmp:
                cnt += 1
            else:
                tmp = i
                cnt = 1
            max_cnt = max(max_cnt, cnt)
        return max_cnt
```