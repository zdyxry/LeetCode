from typing import List

class Solution:
    def smallestRangeI(self, A: List[int], K: int) -> int:
        return max(0, max(A) - min(A) - 2*K)


A = [1]
k = 0
res = Solution().smallestRangeI(A, k)
print(res)