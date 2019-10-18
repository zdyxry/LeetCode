class Solution(object):
    def rangeBitwiseAnd(self, m, n):
        while m < n:
            n &= n -1
        return n

m = 26
n = 30
res = Solution().rangeBitwiseAnd(m, n)
print(res)