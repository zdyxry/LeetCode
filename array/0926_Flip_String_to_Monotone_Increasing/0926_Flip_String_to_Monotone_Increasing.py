
class Solution:
    def minFlipsMonoIncr(self, S: str) -> int:
        N = len(S)
        dp = [0] * 2
        for i in range(1, N+1):
            if S[i-1] == '0':
                dp[0] = dp[0]
                dp[1] = min(dp[1], dp[0]) + 1
            else:
                dp[1] = min(dp[1], dp[0])
                dp[0] = dp[0] + 1
        return min(dp[0], dp[1])


S = "00110"
res = Solution().minFlipsMonoIncr(S)
print(res)