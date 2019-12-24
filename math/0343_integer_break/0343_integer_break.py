
class Solution(object):
    def integerBreak(self, n):
        if n < 4:
            return n - 1
        res = 0
        if n % 3 == 0:
            res = 3 ** (n // 3)
        elif n % 3 == 2:
            res = 3 ** (n // 3) * 2
        else:
            res = 3 ** (n // 3 - 1) * 4
        return res


n = 5
res = Solution().integerBreak(n)
print(res)