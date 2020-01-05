
class Solution(object):
    def myPow(self, x, n):
        result = 1
        abs_n = abs(n)
        while abs_n:
            if abs_n & 1:
                result *= x
            abs_n >>= 1
            x *= x
        return 1 / result if n < 0 else result


x = 2
n = 4
res = Solution().myPow(x, n)
print(res)