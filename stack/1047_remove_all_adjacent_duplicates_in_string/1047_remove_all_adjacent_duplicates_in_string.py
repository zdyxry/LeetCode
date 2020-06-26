class Solution:
    def removeDuplicates(self, S: str) -> str:
        res = []
        for i in S:
            if len(res) >= 1 and i == res[-1]:
                res.pop()
                continue
            res.append(i)
        return ''.join(res)


S = "abbaca"
res = Solution().removeDuplicates(S)
print(res)