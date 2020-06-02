class Solution:
    def hasAllCodes(self, s: str, k: int) -> bool:
        m = set()
        for i in range(len(s)-k+1):
            m.add(s[i:i+k])
        return len(m) == (1 << k)


s = "00110110"
k = 2
res = Solution().hasAllCodes(s, k)
print(res)