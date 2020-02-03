
class Solution(object):
    def getMoneyAmount(self, n):
        dp = [[0] * (n + 1) for _ in range(n + 1)]
        return self.solve(dp, 1, n)

    def solve(self, dp, L, R):
        if L >= R:
            return 0
        if dp[L][R]:
            return dp[L][R]
        dp[L][R] = min(i + max(self.solve(dp, L, i -1), self.solve(dp, i + 1, R)) for i in range(L, R+1))
        return dp[L][R]


n = 1
res = Solution().getMoneyAmount(n)
print(res)