
class Solution(object):
    def closestDivisors(self, num):
        for i in reversed(range(1, int((num + 2) ** 0.5) + 1)):
            if not (num +1) % i:
                return [i, (num+1)//i]
            if not (num+2) % i:
                return [i, (num+2)//i]


num = 8 
res = Solution().closestDivisors(num)
print(res)