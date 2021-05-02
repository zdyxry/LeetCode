class Solution:
    def replaceDigits(self, s: str) -> str:
        m = "abcdefghijklmnopqrstuvwxyz"
        res = ""
        for i in range(len(s)):
            if i % 2 == 1:
                res += m[m.index(s[i-1]) + int(s[i]) ]
            else:
                res += s[i]
            
        return res


s = "a1c1e1"
res = Solution().replaceDigits(s)
print(res)