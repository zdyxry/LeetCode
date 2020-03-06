class Solution(object):
    def tribonacci(self, n):
        self.cache = {}
        def tibo(n):
            if n == 0 :
                return 0 
            if n in (1,2):
                return 1

            if n in self.cache:
                return self.cache[n]
            
            self.cache[n] = tibo(n-1) + tibo(n-2) + tibo(n-3)
            return self.cache[n]
        return tibo(n)


n = 4
res = Solution().tribonacci(n)
print(res)