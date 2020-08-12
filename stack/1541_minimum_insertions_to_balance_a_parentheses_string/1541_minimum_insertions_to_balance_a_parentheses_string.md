1541. Minimum Insertions to Balance a Parentheses String


Medium


Given a parentheses string s containing only the characters '(' and ')'. A parentheses string is balanced if:

```
Any left parenthesis '(' must have a corresponding two consecutive right parenthesis '))'.
Left parenthesis '(' must go before the corresponding two consecutive right parenthesis '))'.
For example, "())", "())(())))" and "(())())))" are balanced, ")()", "()))" and "(()))" are not balanced.
```

You can insert the characters '(' and ')' at any position of the string to balance it if needed.

Return the minimum number of insertions needed to make s balanced.

 

Example 1:

```
Input: s = "(()))"
Output: 1
Explanation: The second '(' has two matching '))', but the first '(' has only ')' matching. We need to to add one more ')' at the end of the string to be "(())))" which is balanced.
```

Example 2:

```
Input: s = "())"
Output: 0
Explanation: The string is already balanced.
```

Example 3:

```
Input: s = "))())("
Output: 3
Explanation: Add '(' to match the first '))', Add '))' to match the last '('.
```

Example 4:

```
Input: s = "(((((("
Output: 12
Explanation: Add 12 ')' to balance the string.
```

Example 5:

```
Input: s = ")))))))"
Output: 5
Explanation: Add 4 '(' at the beginning of the string and one ')' at the end. The string becomes "(((())))))))".
```
 

Constraints:

1 <= s.length <= 10^5  
s consists of '(' and ')' only.


## 方法

```go
func minInsertions(s string) int {
    r := 0
    
    stack := 0
    
    i := 0
    for i < len(s) {
        if s[i] == '(' {
            stack++
            i++
        } else {    // ')'
            i++
            if i == len(s) {
                r++
                if stack == 0 {
                    r++
                } else {
                    stack--
                }
            } else {
                if s[i] != ')' {
                    r++
                } else {
                    i++
                }
                
                if stack == 0 {
                    r++
                } else {
                    stack--
                }
                
            }
        }
    }
    
    return r + stack*2
}
```


```python
class Solution:
    def minInsertions(self, s: str) -> int:
        # use stack
        stack = []
        count = 0
        i=0
        while i < len(s):
            if s[i] == '(':
                stack.append(s[i])
                i += 1
            else:
                # when facing empty stack
                if not stack:
                    if i+1 < len(s) and s[i+1] == ")":
                        count += 1  # add one "("
                        i += 2
                    else:
                        count += 2   # add one "("  and one ")"
                        i += 1
                else:
                    # check two positions
                    if i + 1 < len(s) and s[i+1] == ")":
                        stack.pop()
                        i += 2
                    else:
                        count += 1   # add one ")"
                        stack.pop()
                        i += 1
        
        rest = len(stack)*2  # still have "(" on the stack. one "(" pairs with two ")"
        return count + rest
```