1371. Find the Longest Substring Containing Vowels in Even Counts


Medium


Given the string s, return the size of the longest substring containing each vowel an even number of times. That is, 'a', 'e', 'i', 'o', and 'u' must appear an even number of times.

 

Example 1:

```
Input: s = "eleetminicoworoep"
Output: 13
Explanation: The longest substring is "leetminicowor" which contains two each of the vowels: e, i and o and zero of the vowels: a and u.
```

Example 2:

```
Input: s = "leetcodeisgreat"
Output: 5
Explanation: The longest substring is "leetc" which contains two e's.
```

Example 3:

```
Input: s = "bcbcbc"
Output: 6
Explanation: In this case, the given string "bcbcbc" is the longest because all vowels: a, e, i, o and u appear zero times.
```
 

Constraints:

1 <= s.length <= 5 x 10^5

s contains only lowercase English letters.


## 方法


```go
var target = map[uint8]int{
    'a': 1,
    'e': 2,
    'i': 4,
    'o': 8,
    'u': 16,
}

func findTheLongestSubstring(s string) int {
    var (
        i, curr, rst, st int
        ok               bool
        length           = len(s)
        status           = make([]int, 32)
    )
    for i = 0; i < 32; i++ {
        status[i] = math.MaxInt32
    }
    status[0] = -1
    for i = 0; i < length; i++ {
        if st, ok = target[s[i]]; ok {
            curr ^= st
        }
        if status[curr] == math.MaxInt32 {
            status[curr] = i
        } else {
            rst = max(rst, i-status[curr])
        }
    }
    return rst
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b  
}
```



```python
class Solution(object):
    def findTheLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        vowels = {'a': 1, 'e': 2, 'i': 4, 'o': 8, 'u': 16}
        d, n, r = {0: -1}, 0, 0
        for i, c in enumerate(s):
            if c in vowels:
                n ^= vowels[c]
            print n
            if n not in d:
                d[n] = i
            else:
                r = max(r, i - d[n])
        return r
```