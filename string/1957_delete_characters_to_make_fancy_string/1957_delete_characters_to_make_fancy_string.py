class Solution:
    def makeFancyString(self, s: str) -> str:
        res = []
        for i in s:
            if len(res) >= 2 and i == res[-1] and i == res[-2]:
                continue
            else:
                res.append(i)
        return ''.join(res)
