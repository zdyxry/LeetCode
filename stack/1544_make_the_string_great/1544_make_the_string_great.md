1544. Make The String Great


Easy


Given a string s of lower and upper case English letters.

A good string is a string which doesn't have two adjacent characters s[i] and s[i + 1] where:

```
0 <= i <= s.length - 2
s[i] is a lower-case letter and s[i + 1] is the same letter but in upper-case or vice-versa.
To make the string good, you can choose two adjacent characters that make the string bad and remove them. You can keep doing this until the string becomes good.
```

Return the string after making it good. The answer is guaranteed to be unique under the given constraints.

Notice that an empty string is also good.

 

Example 1:

```
Input: s = "leEeetcode"
Output: "leetcode"
Explanation: In the first step, either you choose i = 1 or i = 2, both will result "leEeetcode" to be reduced to "leetcode".
```

Example 2:

```
Input: s = "abBAcC"
Output: ""
Explanation: We have many possible scenarios, and all lead to the same answer. For example:
"abBAcC" --> "aAcC" --> "cC" --> ""
"abBAcC" --> "abBA" --> "aA" --> ""
```

Example 3:

```
Input: s = "s"
Output: "s"
```

Constraints:

1 <= s.length <= 100  
s contains only lower and upper case English letters.


## 方法


```go
func makeGood(s string) string {
    ans := []byte{}
    for i:=0; i<len(s); i++ {
        ans = append(ans, s[i])
        size := len(ans)
        if size >= 2 {
            // check
            if mequal(ans[size-1], ans[size-2]) {
                ans = ans[:size-2]
            }
        }
    }
    return string(ans)
}

func mequal(a, b byte) bool {
    return (a ^ ' ') == b
}

```


```python
class Solution:
    def makeGood(self, s: str) -> str:
        stack = []
        for i in s:
            if stack and ((ord(stack[-1])-32) == ord(i) or (ord(stack[-1])+32) == ord(i)):
                stack.pop()
            else:
                stack.append(i)
        return ''.join(stack)

```