807. Max Increase to Keep City Skyline


Medium


In a 2 dimensional array grid, each value grid[i][j] represents the height of a building located there. We are allowed to increase the height of any number of buildings, by any amount (the amounts can be different for different buildings). Height 0 is considered to be a building as well. 

At the end, the "skyline" when viewed from all four directions of the grid, i.e. top, bottom, left, and right, must be the same as the skyline of the original grid. A city's skyline is the outer contour of the rectangles formed by all the buildings when viewed from a distance. See the following example.

What is the maximum total sum that the height of the buildings can be increased?

Example:

```
Input: grid = [[3,0,8,4],[2,4,5,7],[9,2,6,3],[0,3,1,0]]
Output: 35
Explanation: 
The grid is:
[ [3, 0, 8, 4], 
  [2, 4, 5, 7],
  [9, 2, 6, 3],
  [0, 3, 1, 0] ]

The skyline viewed from top or bottom is: [9, 4, 8, 7]
The skyline viewed from left or right is: [8, 7, 9, 3]

The grid after increasing the height of buildings without affecting skylines is:

gridNew = [ [8, 4, 8, 7],
            [7, 4, 7, 7],
            [9, 4, 8, 7],
            [3, 3, 3, 3] ]
```

Notes:

1 < grid.length = grid[0].length <= 50.  
All heights grid[i][j] are in the range [0, 100].  
All buildings in grid[i][j] occupy the entire grid cell: that is, they are a 1 x 1 x grid[i][j] rectangular prism.


## 方法

```go
func maxIncreaseKeepingSkyline(grid [][]int) int {
    if grid == nil || len(grid) == 0 || grid[0] == nil ||len(grid[0]) == 0 { return 0 }
    
    maxRow, maxCol := make([]int, len(grid)), make([]int, len(grid[0]))
    rows, cols := len(grid), len(grid[0])
    for i := 0; i < rows; i++ {
        max := 0
        for j := 0; j < cols; j++ {
            max = int(math.Max(float64(max), float64(grid[i][j])))
        }
        maxRow[i] = max
    } 
    for i := 0; i < cols; i++ {
        max := 0
        for j := 0; j < rows; j++ {
            max = int(math.Max(float64(max), float64(grid[j][i])))
        }
        maxCol[i] = max
    }
    res := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            res += (int(math.Min( float64(maxRow[i]), float64(maxCol[j]) )) - grid[i][j])
        }
    }
    return res
}
```



```python
class Solution:
    def maxIncreaseKeepingSkyline(self, grid: List[List[int]]) -> int:
        max_cols = [max(col) for col in zip(*grid)]
        max_rows = [max(row) for row in grid]
        inc = 0
        
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                inc += (min(max_cols[j], max_rows[i]) - grid[i][j])
        return inc
```