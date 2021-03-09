1784. Check if Binary String Has at Most One Segment of Ones


Easy


Given a binary string s ​​​​​without leading zeros, return true​​​ if s contains at most one contiguous segment of ones. Otherwise, return false.

 

Example 1:

```
Input: s = "1001"
Output: false
Explanation: The ones do not form a contiguous segment.
```

Example 2:

```
Input: s = "110"
Output: true
```
 

Constraints:

1 <= s.length <= 100   
s[i]​​​​ is either '0' or '1'.   
s[0] is '1'.   


## 方法


```go
func checkOnesSegment(s string) bool {
    return !strings.Contains(s, "01")
}
```


```python
class Solution:
    def checkOnesSegment(self, s: str) -> bool:
        new = 0
        for i in range(len(s)):
            if s[i] == '0':
                new = 1
            if s[i] == '1':
                if new == 1:
                    return False
        return True
```