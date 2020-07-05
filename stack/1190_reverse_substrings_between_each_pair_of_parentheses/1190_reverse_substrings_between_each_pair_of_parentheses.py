class Solution:
    def reverseParentheses(self, s: str) -> str:
        stack = ['']
        for c in s:
            if c == '(':
                stack.append('')
            elif c == ')':
                reversed_top = stack.pop()[::-1]
                stack[-1] += reversed_top
            else:
                stack[-1] += c
            print(stack)
        return stack[0]


s = "(abcd)"
res = Solution().reverseParentheses(s)
print(res)