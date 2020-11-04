from typing import List
from heapq import heappush, heappop


class Solution:
    def furthestBuilding(self, heights: List[int], bricks: int, ladders: int) -> int:
        heap = []
        for i in range(len(heights) - 1):
            diff = heights[i + 1] - heights[i]
            if diff > 0:
                heappush(heap, diff)
            if len(heap) > ladders:
                bricks -= heappop(heap)
            if bricks < 0:
                return i
        return len(heights) - 1


heights = [4,2,7,6,9,14,12]
bricks = 5
ladders = 1
res = Solution().furthestBuilding(heights, bricks, ladders)
print(res)