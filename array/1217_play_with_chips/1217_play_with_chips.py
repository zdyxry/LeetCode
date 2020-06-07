from typing import List


class Solution:
    def minCostToMoveChips(self, chips: List[int]) -> int:
        even = 0
        for c in chips:
            if c&1==0:
                even+=1
        return min(even,len(chips)-even)


chips = [1,2,3]
res = Solution().minCostToMoveChips(chips)
print(res)