from typing import List

class Solution:
    def findPoisonedDuration(self, timeSeries: List[int], duration: int) -> int:
        n = len(timeSeries)
        if n == 0 :
            return 0

        total = 0
        for i in range(n-1):
            total += min(timeSeries[i+1] - timeSeries[i], duration)
        return total + duration


timeSeries = [1,2]
duration = 2
res = Solution().findPoisonedDuration(timeSeries, duration)
print(res)