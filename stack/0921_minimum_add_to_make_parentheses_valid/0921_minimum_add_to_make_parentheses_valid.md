921. Minimum Add to Make Parentheses Valid


Medium


Given a string S of '(' and ')' parentheses, we add the minimum number of parentheses ( '(' or ')', and in any positions ) so that the resulting parentheses string is valid.

Formally, a parentheses string is valid if and only if:

It is the empty string, or
It can be written as AB (A concatenated with B), where A and B are valid strings, or
It can be written as (A), where A is a valid string.
Given a parentheses string, return the minimum number of parentheses we must add to make the resulting string valid.

 

Example 1:

```
Input: "())"
Output: 1
```

Example 2:

```
Input: "((("
Output: 3
```

Example 3:

```
Input: "()"
Output: 0
```

Example 4:

```
Input: "()))(("
Output: 4
```
 

Note:

S.length <= 1000    
S only consists of '(' and ')' characters.

## 方法

```go
func minAddToMakeValid(S string) int {
    if len(S) == 0 {
		return 0
	}
	stack := make([]rune, 0)
	for _, v := range S {
		if v == '(' {
			stack = append(stack, v)
		} else if (v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return len(stack)
}
```


```go
func minAddToMakeValid(S string) int {
    for strings.Contains(S, "()") {
		S = strings.Replace(S, "()", "", -1)
	}
	return len(S)
}
```



```python
class Solution(object):
    def minAddToMakeValid(self, S):
        """
        :type S: str
        :rtype: int
        """
        add, bal, = 0, 0
        for c in S:
            bal += 1 if c == '(' else -1
            if bal == -1:
                add += 1
                bal += 1
        return add + bal
```