from typing import List


class Solution:
    def minDays(self, bloomDay: List[int], m: int, k: int) -> int:
        n = len(bloomDay)
        if n < m * k:
            return -1
        # 1. 判断第day天是否可以制作出m束花
        def count(day: int):
            num = sums = 0
            for i in range(n):
                # 2. 剪枝
                if num >= m:
                    break
                if bloomDay[i] <= day:
                    sums += 1
                else:
                    sums = 0
                if sums == k:
                    num += 1
                    sums = 0
            return num >= m
        days = sorted(set(bloomDay))
        l, r = 0, len(days)-1
        while l < r:
            mid = l + (r - l) // 2
            if count(days[mid]):
                r = mid
            else:
                l = mid + 1
        return days[l]


bloomDay = [1,10,3,10,2]
m = 3
k = 1
res = Solution().minDays(bloomDay, m, k)
print(res)