1219. Path with Maximum Gold


Medium


In a gold mine grid of size m * n, each cell in this mine has an integer representing the amount of gold in that cell, 0 if it is empty.

Return the maximum amount of gold you can collect under the conditions:

```
Every time you are located in a cell you will collect all the gold in that cell.
From your position you can walk one step to the left, right, up or down.
You can't visit the same cell more than once.
Never visit a cell with 0 gold.
You can start and stop collecting gold from any position in the grid that has some gold.
```

Example 1:

```
Input: grid = [[0,6,0],[5,8,7],[0,9,0]]
Output: 24
Explanation:
[[0,6,0],
 [5,8,7],
 [0,9,0]]
Path to get the maximum gold, 9 -> 8 -> 7.
```

Example 2:

```
Input: grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
Output: 28
Explanation:
[[1,0,7],
 [2,0,6],
 [3,4,5],
 [0,3,0],
 [9,0,20]]
Path to get the maximum gold, 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7.
```

 

Constraints:

1 <= grid.length, grid[i].length <= 15  
0 <= grid[i][j] <= 100  
There are at most 25 cells containing gold.


## 方法

```go
var m, n int

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func dfs(grid [][]int, i, j int) int {
	if grid[i][j] == 0 {
		return 0
	}

	result := 0
	temp := grid[i][j]
	grid[i][j] = 0

	if i > 0 {
		result = max(result, temp+dfs(grid, i-1, j))
	}

	if j > 0 {
		result = max(result, temp+dfs(grid, i, j-1))
	}

	if i < m-1 {
		result = max(result, temp+dfs(grid, i+1, j))
	}

	if j < n-1 {
		result = max(result, temp+dfs(grid, i, j+1))
	}

	grid[i][j] = temp

	return result
}

func getMaximumGold(grid [][]int) int {
	result := 0
	m = len(grid)
	n = len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result = max(result, dfs(grid, i, j))
		}
	}

	return result
}
```

```python
class Solution:
    def helper(self, grid, pos, m, n, count):
        (x, y) = pos
        if x >= 0 and x < m and y >= 0 and y < n and grid[x][y] != 0:
            # print(pos)
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
```