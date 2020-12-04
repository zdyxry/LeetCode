from typing import List

class Solution:
    def minEatingSpeed(self, piles: List[int], H: int) -> int:
        def possible(k):
            # possible <=H
            return sum((p-1)//k+1 for p in piles) <= H
        
        
        l = 1
        h = max(piles)
        while l < h:
            m = (l+h)//2
            if possible(m):
                h = m
            else:
                l = m+1
        return h


piles = [3,6,7,11]
H = 8
res = Solution().minEatingSpeed(piles, H)
print(res)
        