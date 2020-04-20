
class Solution(object):
    def findMinFibonacciNumbers(self, k: int) -> int:
        ls = self.fib(k)
        res = 0
        while k:
            if k >= ls[-1]:
                k -= ls[-1]
                res += 1
            else:
                ls.pop()
        return res

    def fib(self, N) -> int:
        a, b = 0, 1
        res = []
        while b <= N:
            res.append(b)
            a,b = b, a+ b
        return res


k = 7
res = Solution().findMinFibonacciNumbers(k)
print(res)