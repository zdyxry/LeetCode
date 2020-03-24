import math


class Solution(object):
    def sumFourDivisors(self, nums):
        cache = {}
        def fn(n):
            if n not in cache:
                s = set()
                for i in range(1, 1+ int(math.sqrt(n))):
                    if n % i == 0:
                        s.add(i)
                        s.add(n/i)
                    if len(s) > 4:
                        break
                cache[n] = sum(s) if len(s) == 4 else 0
            return cache[n]
        
        return sum(map(fn, nums))


nums = [21,4,7]
res = Solution().sumFourDivisors(nums)
print(res)