from typing import List


class Solution:
    def maxCoins(self, piles: List[int]) -> int:
        # 1,2,2,4,7,8 => 7,2
        # 1,2,3,4,5,6,7,8,9 => 8,6,4
        piles.sort(reverse=True)
        res = []
        i = 0
        while i < len(piles):
            if i % 2:
                i += 1
                continue
            i += 1    
            res.append(piles[i])
            if len(res) == len(piles) /3:
                return sum(res)
        return sum(res)
        


piles = [2,4,1,2,7,8]
res = Solution().maxCoins(piles)
print(res)