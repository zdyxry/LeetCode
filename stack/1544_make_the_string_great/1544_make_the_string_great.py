class Solution:
    def makeGood(self, s: str) -> str:
        stack = []
        for i in s:
            if stack and ((ord(stack[-1])-32) == ord(i) or (ord(stack[-1])+32) == ord(i)):
                stack.pop()
            else:
                stack.append(i)
        return ''.join(stack)


s = "leEeetcode"
res = Solution().makeGood(s)
print(res)