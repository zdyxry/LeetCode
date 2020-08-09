from typing import List

class Solution:
    def findKthPositive(self, arr: List[int], k: int) -> int:
        n = 0
        for i in range(1, arr[-1] + k + 1):
            if i not in arr:
                n += 1
            if n == k:
                return i


arr = [1,2,3,4]
res = Solution().findKthPositive(arr, 2)
print(res)