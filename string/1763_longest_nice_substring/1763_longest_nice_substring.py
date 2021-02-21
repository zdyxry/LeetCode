class Solution:
    def longestNiceSubstring(self, s: str) -> str:
        res = ''
        n = len(s)
        for i in range(n):
            for j in range(i+1, n):
                t = s[i:j+1]
                if all(c.upper() in t and c.lower() in t for c in t):
                    if len(t) > len(res):
                        res = t
        return res


s = "YazaAay"
res = Solution().longestNiceSubstring(s)
print(res)