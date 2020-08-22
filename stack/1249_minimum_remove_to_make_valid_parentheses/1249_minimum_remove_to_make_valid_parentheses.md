1249. Minimum Remove to Make Valid Parentheses


Medium


Given a string s of '(' , ')' and lowercase English characters. 

Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions ) so that the resulting parentheses string is valid and return any valid string.

Formally, a parentheses string is valid if and only if:

```
It is the empty string, contains only lowercase characters, or
It can be written as AB (A concatenated with B), where A and B are valid strings, or
It can be written as (A), where A is a valid string.
```

Example 1:

```
Input: s = "lee(t(c)o)de)"
Output: "lee(t(c)o)de"
Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
```

Example 2:

```
Input: s = "a)b(c)d"
Output: "ab(c)d"
```

Example 3:

```
Input: s = "))(("
Output: ""
Explanation: An empty string is also valid.
```

Example 4:

```
Input: s = "(a(b(c)d)"
Output: "a(b(c)d)"
```
 

Constraints:

1 <= s.length <= 10^5  
s[i] is one of  '(' , ')' and lowercase English letters.


## 方法

```go
func minRemoveToMakeValid(s string) string {
    stack := make([]int,0,len(s))

    toDelete := make([]int,len(s))

    for i:=0;i<len(s);i++{
    	if s[i] == '(' {
    		stack = append(stack,i)
    	} else if s[i] == ')' {
    		if len(stack) == 0 {
    			toDelete[i] = 1
    		} else {
    			stack = stack[:len(stack)-1]
    		}
    	}
    }

    for i:=0; i<len(stack);i++{
    	toDelete[stack[i]] = 1
    }

    var sb strings.Builder

    for i:=0;i<len(s);i++ {
    	if toDelete[i] == 0 {
    		sb.WriteByte(s[i])
    	}
    }

    return sb.String()
}
```


```python
class Solution:
    def minRemoveToMakeValid(self, s: str) -> str:
        indices, ans = [], list(s)
        for i, c in enumerate(s):
            if c == ')':
                if indices:
                    indices.pop()
                else:
                    ans[i] = ''
            elif c == '(':
                indices.append(i)

        for i in indices:
            ans[i] = ''

        return ''.join(ans)
```


```python
class Solution:
    def minRemoveToMakeValid(self, s: str) -> str:
        stack = []
        res = []
        for idx, i in enumerate(s):
            if i in ['(', ')']:
                if i == '(':
                    stack.append(i)
                    res.append(idx)
                elif i == ')':
                    if len(stack) > 0 and stack[-1] == '(':
                        stack.pop()
                        res.pop()
                    else:
                        stack.append(i)
                        res.append(idx)
        print(stack)
        print(res)
        tmp = []
        for idx, i in enumerate(s):
            if idx in res:
                continue
            else:
                tmp.append(i)
        return ''.join(tmp)
```