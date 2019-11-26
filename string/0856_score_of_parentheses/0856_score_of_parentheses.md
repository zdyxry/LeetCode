856. Score of Parentheses


Medium


Given a balanced parentheses string S, compute the score of the string based on the following rule:

* () has score 1  
* AB has score A + B, where A and B are balanced parentheses strings.  
* (A) has score 2 * A, where A is a balanced parentheses string.  
 

Example 1:

```
Input: "()"
Output: 1
```

Example 2:

```
Input: "(())"
Output: 2
```

Example 3:

```
Input: "()()"
Output: 2
```

Example 4:

```
Input: "(()(()))"
Output: 6
```
 

Note:

S is a balanced parentheses string, containing only ( and ).  
2 <= S.length <= 50  


## 方法



```go
func scoreOfParentheses(s string) int {
    res := 0
	factor := 1
	size := len(s)
	for i := 0; i < size; i++ {
		if s[i] == '(' {
			factor *= 2
		} else {
			factor /= 2
		}
		if s[i] == '(' && s[i+1] == ')' {
			res += factor / 2
		}
	}
	return res
}
```



```python
class Solution(object):
    def scoreOfParentheses(self, S):
        """
        :type S: str
        :rtype: int
        """
        result, depth = 0, 0
        for i in xrange(len(S)):
            if S[i] == '(':
                depth += 1
            else:
                depth -= 1
                if S[i-1] == '(':
                    result += 2**depth
        return result
```