

class Solution:
    def xorOperation(self, n: int, start: int) -> int:
        ans = 0
        for i in range(n):
            ans ^= start + 2 * i
        return ans


n = 5
start = 0
res = Solution().xorOperation(n, start)
print(res)