class Solution:
    def nthPersonGetsNthSeat(self, n: int) -> float:
        return 1 if n==1 else .5



n = 1
res = Solution().nthPersonGetsNthSeat(n)
print(res)