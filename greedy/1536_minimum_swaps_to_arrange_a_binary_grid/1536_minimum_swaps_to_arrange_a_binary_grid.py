from typing import List

class Solution:
    def minSwaps(self, grid: List[List[int]]) -> int:
        need = 0
        l = len(grid)
        for i in range(l-1):
            for j in range(i, l):
                if sum(grid[j][i+1:]) == 0:
                    need += j-i
                    grid[i:j+1] = grid[j:j+1]+grid[i:j]
                    break
            else:
                return -1
        return need


grid = [[0,0,1],[1,1,0],[1,0,0]]
res = Solution().minSwaps(grid)
print(res)