from typing import List


class Solution:
    def getNoZeroIntegers(self, n: int) -> List[int]:
        for i in range(1, n):
            if self.check(i) and self.check(n - i):
                return [i, n - i]

    def check(self, n: int) -> bool:
        while n > 0:
            if n % 10 == 0:
                return False
            n //= 10
        return True


n = 2
res = Solution().getNoZeroIntegers(n)
print(res)