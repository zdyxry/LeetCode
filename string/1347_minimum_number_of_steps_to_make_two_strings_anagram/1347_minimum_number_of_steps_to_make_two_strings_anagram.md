1347. Minimum Number of Steps to Make Two Strings Anagram


Medium


Given two equal-size strings s and t. In one step you can choose any character of t and replace it with another character.

Return the minimum number of steps to make t an anagram of s.

An Anagram of a string is a string that contains the same characters with a different (or the same) ordering.

 

Example 1:

```
Input: s = "bab", t = "aba"
Output: 1
Explanation: Replace the first 'a' in t with b, t = "bba" which is anagram of s.
```

Example 2:

```
Input: s = "leetcode", t = "practice"
Output: 5
Explanation: Replace 'p', 'r', 'a', 'i' and 'c' from t with proper characters to make t anagram of s.
```

Example 3:

```
Input: s = "anagram", t = "mangaar"
Output: 0
Explanation: "anagram" and "mangaar" are anagrams. 
```

Example 4:

```
Input: s = "xxyyzz", t = "xxyyzz"
Output: 0
```

Example 5:

```
Input: s = "friend", t = "family"
Output: 4
```
 

Constraints:

1 <= s.length <= 50000

s.length == t.length

s and t contain lower-case English letters only.

## 方法

```go
func minSteps(s string, t string) int {
    var commonChar = 0
    var charA = map[uint8]int{}

    for i := 0; i < len(s); i++ {
        charA[s[i]]++
    }
    for i := 0; i < len(t); i++ {
        charA[t[i]]--
    }
    for _, v := range charA {
        commonChar += abs(v)
    }
    return commonChar / 2
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
```


```python
class Solution(object):
    def minSteps(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: int
        """
        cnt = [0] * 26
        for i in range(len(s)):
            cnt[ord(s[i]) - ord('a')] += 1
            cnt[ord(t[i]) - ord('a')] -= 1
        res = 0
        for i in cnt:
            if i > 0:
                res += i
        return res
```