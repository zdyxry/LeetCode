from typing import List

class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        def dfs(gird, i, j):
            if 0<=i<m and  0<=j<n and grid[i][j]:
                grid[i][j] = 0
                return 1 + dfs(grid, i+1,j) + dfs(grid, i-1, j) + dfs(grid, i, j+1) + dfs(grid, i, j-1)
            return 0
        result = 0
        for x in range(m):
            for y in range(n):
                result = max(result, dfs(grid, x, y))
        return result


grid = [[1,1,0,0,0],[1,1,0,0,0],[0,0,0,1,1],[0,0,0,1,1]]
res = Solution().maxAreaOfIsland(grid)
print(res)