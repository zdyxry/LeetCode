1941. Check if All Characters Have Equal Number of Occurrences


Easy


Given a string s, return true if s is a good string, or false otherwise.

A string s is good if all the characters that appear in s have the same number of occurrences (i.e., the same frequency).

 

Example 1:

```
Input: s = "abacbc"
Output: true
Explanation: The characters that appear in s are 'a', 'b', and 'c'. All characters occur 2 times in s.
```

Example 2:

```
Input: s = "aaabb"
Output: false
Explanation: The characters that appear in s are 'a' and 'b'.
'a' occurs 3 times while 'b' occurs 2 times, which is not the same number of times.
```

Constraints:

1 <= s.length <= 1000   
s consists of lowercase English letters.


## 方法


```go
func areOccurrencesEqual(s string) bool {
    m := make(map[rune]int)
    for _, c := range s {
        m[c]++
    }
    
    current := 0
    for _, v := range m {
        if current == 0 {
            current = v
        } else {
            if current != v {
                return false
            }
        }
    }
    return true
}
```


```
class Solution:
    def areOccurrencesEqual(self, s: str) -> bool:
        m = {}
        for c in s:
            if c in m:
                m[c] += 1
            else:
                m[c] = 0
        return len(set([x for x in m.values()])) == 1
```