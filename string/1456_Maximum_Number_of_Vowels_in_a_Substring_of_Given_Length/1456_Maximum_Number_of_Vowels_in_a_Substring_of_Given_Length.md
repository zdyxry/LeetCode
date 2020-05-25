1456. Maximum Number of Vowels in a Substring of Given Length


Medium

Given a string s and an integer k.

Return the maximum number of vowel letters in any substring of s with length k.

Vowel letters in English are (a, e, i, o, u).

 

Example 1:

```
Input: s = "abciiidef", k = 3
Output: 3
Explanation: The substring "iii" contains 3 vowel letters.
```

Example 2:

```
Input: s = "aeiou", k = 2
Output: 2
Explanation: Any substring of length 2 contains 2 vowels.
```

Example 3:

```
Input: s = "leetcode", k = 3
Output: 2
Explanation: "lee", "eet" and "ode" contain 2 vowels.
```

Example 4:

```
Input: s = "rhythms", k = 4
Output: 0
Explanation: We can see that s doesn't have any vowel letters.
```

Example 5:

```
Input: s = "tryhard", k = 4
Output: 1
```
 

Constraints:
  
1 <= s.length <= 10^5  
s consists of lowercase English letters.  
1 <= k <= s.length

## 方法


```go
func maxVowels(s string, k int) int {
    res := 0
    dp := make([]int,len(s)+1)
    count := 0
    for i:=0;i<len(s);i++{
        dp[i] = count
        if check(s[i]) == true{
            count++
        }
    }
    dp[len(s)] = count
    for i:=k;i<=len(s);i++{
        if dp[i]-dp[i-k] > res{
            res = dp[i]-dp[i-k]
        }
    }
    return res
}

func check(str byte) bool{
    if str == 'a' || str == 'e' || str == 'i' || str == 'o' || str == 'u'{
        return true
    }
    return false
}
```


```python
class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        d = {'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}
        for c in s[:k]:
            if c in d:
                d[c] += 1
        res = sum(d.values())
        for i in range(k, len(s)):
            if s[i-k] in d:
                d[s[i-k]] -= 1
            if s[i] in d:
                d[s[i]] += 1
            res = max(res, sum(d.values()))
        return res
```