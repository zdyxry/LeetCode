from typing import List
import collections


class Solution:
    def numIdenticalPairs(self, A: List[int]) -> int:
        return sum(k * (k - 1) / 2 for k in collections.Counter(A).values())


A = [1,2,3,1,1,3]
res = Solution().numIdenticalPairs(A)
print(res)