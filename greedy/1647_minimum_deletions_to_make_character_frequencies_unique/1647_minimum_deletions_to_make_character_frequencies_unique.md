1647. Minimum Deletions to Make Character Frequencies Unique


Medium


A string s is called good if there are no two different characters in s that have the same frequency.

Given a string s, return the minimum number of characters you need to delete to make s good.

The frequency of a character in a string is the number of times it appears in the string. For example, in the string "aab", the frequency of 'a' is 2, while the frequency of 'b' is 1.

 

Example 1:

```
Input: s = "aab"
Output: 0
Explanation: s is already good.
```

Example 2:

```
Input: s = "aaabbbcc"
Output: 2
Explanation: You can delete two 'b's resulting in the good string "aaabcc".
Another way it to delete one 'b' and one 'c' resulting in the good string "aaabbc".
```

Example 3:

```
Input: s = "ceabaacb"
Output: 2
Explanation: You can delete both 'c's resulting in the good string "eabaab".
Note that we only care about characters that are still in the string at the end (i.e. frequency of 0 is ignored).
```
 

Constraints:
 
1 <= s.length <= 105   
s contains only lowercase English letters.


## 方法



```go
// count
func minDeletions(s string) int {

    counts := make([]int, 26)
    for i := range s {
        counts[s[i]-'a']++
    }

    sort.Ints(counts)

    res := 0
    for i := len(counts)-2; i > -1; i-- {
        if counts[i] > 0 && counts[i] >= counts[i+1] {
            target := max(0, counts[i+1]-1)
            res += (counts[i] - target)
            counts[i] = target
        }
    }

    return res
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
    def minDeletions(self, s: str) -> int:
        c = collections.Counter(s)
        seen = set()
        res = 0
        for k, v in c.items():
            cur = v
            while cur>0 and cur in seen:
                cur -= 1
            res += v - cur
            if cur>0:
                seen.add(cur)
        return res
```