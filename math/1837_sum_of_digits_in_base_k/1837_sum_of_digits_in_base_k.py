class Solution:
    def sumBase(self, n: int, k: int) -> int:
        res = 0
        while n:
            res += n % k
            n //= k
        return res


n = 34
k = 6
res = Solution().sumBase(n, k)
print(res)