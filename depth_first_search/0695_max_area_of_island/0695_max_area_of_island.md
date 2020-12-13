695. Max Area of Island


Medium

Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

Example 1:

```
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.
```

Example 2:

```
[[0,0,0,0,0,0,0,0]]
Given the above grid, return 0.
Note: The length of each dimension in the given grid does not exceed 50.
```


## 方法

```go
func maxAreaOfIsland(grid [][]int) int {
    if len(grid) == 0{return 0}

    var res int // default 0
    
    row, col := len(grid), len(grid[0])

    for x := 0; x < row; x++{
        for y := 0; y < col; y++{
            if grid[x][y] == 1{
                res = Max(res, dfs(x, y, grid))
            }
        }
    }
    return res
}

func Max(a, b int) int{
    if a > b{
        return a
    }
    return b
}

func dfs(x, y int, grid [][]int) int{
    if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0{
        return 0
    }
    grid[x][y] = 0

    count := 1
    count += dfs(x - 1, y, grid)
    count += dfs(x + 1, y, grid)
    count += dfs(x, y - 1 , grid)
    count += dfs(x, y + 1, grid)
    return count
}

```

```python
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
```