class Solution:
    def divisorGame(self, N: int) -> bool:
        return not N%2


N = 2
res = Solution().divisorGame(N)
print(res)