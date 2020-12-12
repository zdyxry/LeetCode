from typing import List

class Solution:
    def projectionArea(self, grid: List[List[int]]) -> int:
        xy = sum(grid[i][j]>0  for i in range(len(grid)) for j in range(len(grid[0])))
        xz = sum(max(i)  for i in grid)
        yz = sum(max(i)  for i in zip(*grid))

        return xy+xz+yz



grid = [[1,2],[3,4]]
res = Solution().projectionArea(grid)
print(res)