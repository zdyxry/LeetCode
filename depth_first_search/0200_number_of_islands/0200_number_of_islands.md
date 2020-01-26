200. Number of Islands

Medium

Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:

```
Input:
11110
11010
11000
00000

Output: 1
```

Example 2:

```
Input:
11000
11000
00100
00011

Output: 3
```


## 方法



```go
func numIslands(grid [][]byte) int {
    if len(grid) == 0 {
        return 0
    }
    lenRow, lenCol := len(grid), len(grid[0])
    res := 0
    for r := 0; r < lenRow; r++ {
        for c := 0; c < lenCol; c++ {
            if grid[r][c] == '1' {
                res++
                dfs(&grid, r, c)
            }
        }
    }
    return res
}

func dfs(grid *[][]byte, r, c int) {
    lenRow, lenCol := len(*grid), len((*grid)[0])
    if r < 0 || r >= lenRow || c < 0 || c >= lenCol || (*grid)[r][c] == '0' {
        return
    }
    (*grid)[r][c] = '0'
    dfs(grid, r - 1, c)
    dfs(grid, r + 1, c)
    dfs(grid, r, c - 1)
    dfs(grid, r, c + 1)
}
```



```python
class Solution(object):
    def numIslands(self, grid):
        """
        :type grid: List[List[str]]
        :rtype: int
        """
        if grid is None or len(grid) is 0:
            return 0
        numIslands = 0
        for y in range(len(grid)):
            for x in range(len(grid[y])):
                if (grid[y][x] == "1"):
                    numIslands += self.dfs(grid, y, x)
        return numIslands
    
    def dfs(self, grid, y, x):
        if  ((y<0) or (y>=len(grid)) or (x<0) or (x>=len(grid[y])) or (grid[y][x] == "0")):
            return 0
        grid[y][x] = "0"
        self.dfs(grid, y-1, x,)
        self.dfs(grid, y+1, x,)
        self.dfs(grid, y, x-1,)
        self.dfs(grid, y, x+1)
        return 1
```