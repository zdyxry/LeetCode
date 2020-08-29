1021. Remove Outermost Parentheses


Easy


A valid parentheses string is either empty (""), "(" + A + ")", or A + B, where A and B are valid parentheses strings, and + represents string concatenation.  For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.

A valid parentheses string S is primitive if it is nonempty, and there does not exist a way to split it into S = A+B, with A and B nonempty valid parentheses strings.

Given a valid parentheses string S, consider its primitive decomposition: S = P_1 + P_2 + ... + P_k, where P_i are primitive valid parentheses strings.

Return S after removing the outermost parentheses of every primitive string in the primitive decomposition of S.

 

Example 1:

```
Input: "(()())(())"
Output: "()()()"
Explanation: 
The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
After removing outer parentheses of each part, this is "()()" + "()" = "()()()".
```

Example 2:

```
Input: "(()())(())(()(()))"
Output: "()()()()(())"
Explanation: 
The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".
```

Example 3:

```
Input: "()()"
Output: ""
Explanation: 
The input string is "()()", with primitive decomposition "()" + "()".
After removing outer parentheses of each part, this is "" + "" = "".
```

Note:

S.length <= 10000  
S[i] is "(" or ")"  
S is a valid parentheses string


## 方法

```go
func removeOuterParentheses(S string) string {
    if len(S)==0{
	  	return ""
	  }
	  stack := []byte{S[0]}
	  start :=1
	  var res strings.Builder
      for i:=1;i<len(S);i++{
		  if len(stack)>0&&S[i]==')'&&stack[len(stack)-1]=='(' {
		  	  stack = stack[:len(stack)-1]
		  	  if len(stack)==0{
		  	  	  res.WriteString(S[start:i])
				  start = i+2
			  }
		  }else {
		  	 stack = append(stack,S[i])
		  }
	  }
	  return res.String()
}
```

```python
class Solution:
    def removeOuterParentheses(self, S: str) -> str:
        stack = []
        i = 0
        res = ''
        while i < len(S):
            if not stack:
                stack.append(S[i])
            else:
                if S[i] == stack[-1]:
                    if len(stack) >= 1:
                        res += S[i]
                    stack.append(S[i])
                else:
                    if len(stack) > 1:
                        res += S[i]
                    stack.pop()
            i += 1
        return res
```