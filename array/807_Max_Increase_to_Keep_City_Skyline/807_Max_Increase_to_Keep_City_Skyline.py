from typing import List

class Solution:
    def maxIncreaseKeepingSkyline(self, grid: List[List[int]]) -> int:
        max_cols = [max(col) for col in zip(*grid)]
        max_rows = [max(row) for row in grid]
        inc = 0
        
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                inc += (min(max_cols[j], max_rows[i]) - grid[i][j])
        return inc


grid = [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]
res = Solution().maxIncreaseKeepingSkyline(grid)
print(res)