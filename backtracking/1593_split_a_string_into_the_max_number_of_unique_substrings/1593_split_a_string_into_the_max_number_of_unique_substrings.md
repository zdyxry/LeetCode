1593. Split a String Into the Max Number of Unique Substrings


Medium


Given a string s, return the maximum number of unique substrings that the given string can be split into.

You can split string s into any list of non-empty substrings, where the concatenation of the substrings forms the original string. However, you must split the substrings such that all of them are unique.

A substring is a contiguous sequence of characters within a string.

 

Example 1:

```
Input: s = "ababccc"
Output: 5
Explanation: One way to split maximally is ['a', 'b', 'ab', 'c', 'cc']. Splitting like ['a', 'b', 'a', 'b', 'c', 'cc'] is not valid as you have 'a' and 'b' multiple times.
```

Example 2:

```
Input: s = "aba"
Output: 2
Explanation: One way to split maximally is ['a', 'ba'].
```

Example 3:

```
Input: s = "aa"
Output: 1
Explanation: It is impossible to split the string any further.
```

Constraints:

1 <= s.length <= 16

s contains only lower case English letters.


## 方法


```go

func maxUniqueSplit(s string) int {
    return helper(s, map[string]bool{})
}

func helper(s string, seen map[string]bool) int {
    ans := 0
    for i := 1; i <= len(s); i++ {
        if !seen[s[0:i]] {
            seen[s[0:i]] = true
            ans = max(ans, 1 + helper(s[i:], seen))
            seen[s[0:i]] = false
        }
    }
    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```


```python
class Solution:
    def maxUniqueSplit(self, s: str) -> int:
        self.ans = 0
        d = set()

        def helper(num, sub):
            if not sub:
                self.ans = max(num, self.ans)
                return
            
            if sub in d: return 
            
            for i in range(1, len(sub) + 1):
                if sub[:i] not in d:
                    d.add(sub[:i])
                    helper(num + 1, sub[i:])
                    d.remove(sub[:i]) 
                    
        helper(0, s)

        return self.ans
```