1415. The k-th Lexicographical String of All Happy Strings of Length n


Medium

A happy string is a string that:

consists only of letters of the set ['a', 'b', 'c'].  
s[i] != s[i + 1] for all values of i from 1 to s.length - 1 (string is 1-indexed).  
For example, strings "abc", "ac", "b" and "abcbabcbcb" are all happy strings and strings "aa", "baa" and "ababbc" are not happy strings.  

Given two integers n and k, consider a list of all happy strings of length n sorted in lexicographical order.  

Return the kth string of this list or return an empty string if there are less than k happy strings of length n.

 

Example 1:

```
Input: n = 1, k = 3
Output: "c"
Explanation: The list ["a", "b", "c"] contains all happy strings of length 1. The third string is "c".
```

Example 2:

```
Input: n = 1, k = 4
Output: ""
Explanation: There are only 3 happy strings of length 1.
```

Example 3:

```
Input: n = 3, k = 9
Output: "cab"
Explanation: There are 12 different happy string of length 3 ["aba", "abc", "aca", "acb", "bab", "bac", "bca", "bcb", "cab", "cac", "cba", "cbc"]. You will find the 9th string = "cab"
```

Example 4:

```
Input: n = 2, k = 7
Output: ""
```

Example 5:

```
Input: n = 10, k = 100
Output: "abacbabacb"
```

Constraints:

1 <= n <= 10  
1 <= k <= 100


## 方法

```go
func getHappyString(n int, k int) string {
    res := ""
    num := 0
    var dfs func(i int, buf []byte)
    dfs = func(i int, buf []byte) {
        if i == n {
            num++
            if num == k {
                res = string(buf)
            }
            return
        }
        for j := byte('a'); j < 'd'; j++ {
            if i == 0 || j != buf[i-1] {
                buf[i] = j
                dfs(i+1, buf)
            }
        }
    }
    dfs(0, make([]byte, n))
    return res
}
```



```python
class Solution:
    def getHappyString(self, n: int, k: int) -> str:
        self.res = ""
        self.n = n
        self.k = k
        def dfs(cur, tmp):
            if len(cur) == self.n:
                self.k -= 1
                if self.k == 0:
                    self.res = cur
                    return
                return
            for i in tmp:
                if self.res:
                    return self.res
                new_tmp = [j for j in ["a", "b", "c"] if j != i]
                dfs(cur+i, new_tmp)
        dfs("", ["a", "b", "c"])
        return self.res
```