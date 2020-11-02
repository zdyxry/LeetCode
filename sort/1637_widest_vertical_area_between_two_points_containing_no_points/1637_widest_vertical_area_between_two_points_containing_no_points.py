from typing import List

class Solution:
    def maxWidthOfVerticalArea(self, points: List[List[int]]) -> int:
        d = sorted({x for x, y in points})
        return max(d[i] - d[i-1] for i in range(1, len(d)))


points = [[8,7],[9,9],[7,4],[9,7]]
res = Solution().maxWidthOfVerticalArea(points)
print(res)