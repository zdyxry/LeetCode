class Solution:
    def licenseKeyFormatting(self, S: str, K: int) -> str:
        
        S = S.upper().replace('-','')
        size = len(S)
        s1 = K if size%K==0 else size%K
        res = S[:s1]
        while s1<size:
            res += '-'+S[s1:s1+K]
            s1 += K
        return res


S = "5F3Z-2e-9-w"
K = 4
res = Solution().licenseKeyFormatting(S, K)
print(res)