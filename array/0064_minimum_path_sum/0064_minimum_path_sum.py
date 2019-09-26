class Solution(object):
    def minPathSum(self, grid):
        sum = list(grid[0])
        for j in xrange(1, len(grid[0])):
            sum[j] = sum[j-1] + grid[0][j]

        for i in xrange(1, len(grid)):
            sum[0] += grid[i][0]
            for j in xrange(1, len(grid[0])):
                sum[j] = min(sum[j-1], sum[j]) + grid[i][j]
            
        return sum[-1]

grid = [[1,3,1],[1,5,1],[4,2,1]]
res = Solution().minPathSum(grid)
print(res)