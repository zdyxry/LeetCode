class Solution:
    def rotatedDigits(self, N: int) -> int:
        ans = 0
        # For each x in [1, N], check whether it's good
        for x in range(1, N+1):
            S = str(x)
            # Each x has only rotateable digits, and one of them
            # rotates to a different digit
            ans += (all(d not in '347' for d in S)
                    and any(d in '2569' for d in S))
        return ans


N = 10
res = Solution().rotatedDigits(N)
print(res)