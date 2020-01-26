
class Solution(object):
    def numIslands(self, grid):
        if grid is None or len(grid) is 0:
            return 0
        numIslands = 0
        for y in range(len(grid)):
            for x in range(len(grid[y])):
                if (grid[y][x] == "1"):
                    numIslands += self.dfs(grid, y, x)
        return numIslands

    def dfs(self, grid, y, x):
        if ((y < 0) or (y >= len(grid)) or (x < 0) or (x > len(grid[y])) or (grid[y][x] == "0")):
            return 0
        grid[y][x] = "0"
        self.dfs(grid, y - 1, x,)
        self.dfs(grid, y + 1, x,)
        self.dfs(grid, y, x - 1)
        self.dfs(grid, y, x + 1)
        return 1


grid = [["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]

res = Solution().numIslands(grid)
print(res)
    