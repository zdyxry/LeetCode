1653. Minimum Deletions to Make String Balanced


Medium


You are given a string s consisting only of characters 'a' and 'b'​​​​.

You can delete any number of characters in s to make s balanced. s is balanced if there is no pair of indices (i,j) such that i < j and s[i] = 'b' and s[j]= 'a'.

Return the minimum number of deletions needed to make s balanced.

 

Example 1:

```
Input: s = "aababbab"
Output: 2
Explanation: You can either:
Delete the characters at 0-indexed positions 2 and 6 ("aababbab" -> "aaabbb"), or
Delete the characters at 0-indexed positions 3 and 6 ("aababbab" -> "aabbbb").
```

Example 2:

```
Input: s = "bbaaaaabb"
Output: 2
Explanation: The only solution is to delete the first two characters.
```

Constraints:

1 <= s.length <= 105   
s[i] is 'a' or 'b'​​.


## 方法


```go
func minimumDeletions(s string) int {
    n := len(s)
    a := make([]int, n, n)
    b := make([]int, n, n)
    
    if s[0] == 'a' {a[0] = 1}
    if s[n - 1] == 'b' {b[n - 1] = 1}
    
    for i := 1; i < n; i ++ {
        a[i] = a[i - 1]
        if s[i] == 'a' {a[i] ++}
    }
    
    for j := n - 2; j >= 0; j -- {
        b[j] = b[j + 1]
        if s[j] == 'b' {b[j] ++}
    }
    // fmt.Println(a, b)
    res := 0
    
    for i := 0; i < n; i ++ {
        x := a[i] + b[i]
        if x > res {
            res = x
        }
    }
    
    return n - res
}

```


```python
class Solution:
    def minimumDeletions(self, s: str) -> int:
        n = len(s)
        a = 0
        res = 0
        for i in range(n - 1, -1, -1):
            if s[i] == 'a':
                a += 1
            elif s[i] == 'b':
                if a > 0:
                    a -= 1
                    res += 1
        return res
```