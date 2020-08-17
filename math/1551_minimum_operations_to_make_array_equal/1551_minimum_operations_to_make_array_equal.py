class Solution:
    def minOperations(self, n: int) -> int:
        return n*n // 4


n = 4
res = Solution().minOperations(n)
print(res)