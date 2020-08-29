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


S = "(()())(())"
res = Solution().removeOuterParentheses(S)
print(res)