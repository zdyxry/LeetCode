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


s = "lee(t(c)o)de)"
res = Solution().minRemoveToMakeValid(s)
print(res)