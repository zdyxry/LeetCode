class Solution:
    def bitwiseComplement(self, N: int) -> int:
        Nbin=bin(N)
        return 2**len(Nbin[2:])-1-N


N = 5
res = Solution().bitwiseComplement(N)
print(res)