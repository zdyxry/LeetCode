20. Valid Parentheses


Easy


Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:
```
Input: "()"
Output: true
```
Example 2:
```
Input: "()[]{}"
Output: true
```
Example 3:
```
Input: "(]"
Output: false
```
Example 4:
```
Input: "([)]"
Output: false
```
Example 5:
```
Input: "{[]}"
Output: true
```


# 方法

采用 stack 方式解决，当检测到括号的另一边，则对应栈顶一定与其匹配。


```go
func isValid(s string) bool {
    size := len(s)

	stack := make([]byte, size)
	top := 0

	for i := 0; i < size; i++ {
		c := s[i]
		switch c {
		case '(':
			stack[top] = c + 1 // '('+1 is ')'
			top++
		case '[', '{':
			stack[top] = c + 2
			top++
		case ')', ']', '}':
            fmt.Println(stack, top)
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}
		}
	}

	return top == 0
}
```


```python
class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        stack, lookup = [], {"(": ")", "{": "}", "[": "]"}
        for parenthese in s:
            if parenthese in lookup:
                stack.append(parenthese)
            elif len(stack) == 0 or lookup[stack.pop()] != parenthese:
                return False
        return len(stack) == 0
```