from typing import List


class Solution:
    def helper(self, grid, pos, m, n, count):
        x,y = pos
        if x >= 0 and x < m and y >= 0 and y < n and grid[x][y] != 0:
            count += grid[x][y]
            temp = grid[x][y]
            grid[x][y] = 0
            self.helper(grid, (x-1, y), m, n, count)
            self.helper(grid, (x+1, y), m, n, count)
            self.helper(grid, (x, y-1), m, n, count)
            self.helper(grid, (x, y+1), m, n, count)
            grid[x][y] = temp
        self.res.append(count)
    
    def getMaximumGold(self, grid: List[List[int]]) -> int:
        self.res = []
        m = len(grid)
        n = len(grid[0])
        for i in range(m):
            for j in range(n):
                self.helper(grid, (i, j), m, n, 0)
        return max(self.res)


grid = [[0,6,0],[5,8,7],[0,9,0]]
res = Solution().getMaximumGold(grid)
print(res)