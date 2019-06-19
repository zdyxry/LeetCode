463. Island Perimeter

Easy

You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.

 

Example:

```
Input:
[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]

Output: 16

Explanation: The perimeter is the 16 yellow stripes in the image below:

```


# 方法
遍历二维数组，当当前值为1时，周长为4，考虑当前值上、下、左、右是否为1，若为1，则周长减1。


```go
func islandPerimeter(grid [][]int) int {
	var dx = []int{-1, 1, 0, 0}
	var dy = []int{0, 0, -1, 1}
	m, n := len(grid), len(grid[0])

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			res += 4
			for k := 0; k < 4; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1 {
					res--
				}
			}
		}
	}

	return res
}
```


```python
class Solution:
    def islandPerimeter(self, grid: List[List[int]]) -> int:
        res = 0
        for row in range(len(grid)):
            for column in range(len(grid[0])):
                    if grid[row][column]:
                        res+=4
                        if row and grid[row-1][column]:
                            res-=2                        
                        if column and grid[row][column-1]:
                            res-=2
        return res
```


