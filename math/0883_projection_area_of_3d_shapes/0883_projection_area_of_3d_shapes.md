883. Projection Area of 3D Shapes


Easy


On a N * N grid, we place some 1 * 1 * 1 cubes that are axis-aligned with the x, y, and z axes.

Each value v = grid[i][j] represents a tower of v cubes placed on top of grid cell (i, j).

Now we view the projection of these cubes onto the xy, yz, and zx planes.

A projection is like a shadow, that maps our 3 dimensional figure to a 2 dimensional plane. 

Here, we are viewing the "shadow" when looking at the cubes from the top, the front, and the side.

Return the total area of all three projections.

 

Example 1:

```
Input: [[2]]
Output: 5
```

Example 2:

```
Input: [[1,2],[3,4]]
Output: 17
Explanation: 
Here are the three projections ("shadows") of the shape made with each axis-aligned plane.
```

Example 3:

```
Input: [[1,0],[0,2]]
Output: 8
```

Example 4:

```
Input: [[1,1,1],[1,0,1],[1,1,1]]
Output: 14
```

Example 5:

```
Input: [[2,2,2],[2,1,2],[2,2,2]]
Output: 21
```
 

Note:

1 <= grid.length = grid[0].length <= 50   
0 <= grid[i][j] <= 50


## 方法


```go
func projectionArea(grid [][]int) int {
	rowNum, colNum := len(grid), len(grid[0])
	xy, xz, yz, rowMax, colMax := 0, 0, 0, make([]int, rowNum), make([]int, colNum)
	for row, v1 := range grid {
		for col, v2 := range v1 {
			if v2 != 0 {
				xy++
			}
			rowMax[row] = int(math.Max(float64(rowMax[row]), float64(grid[row][col])))
			colMax[col] = int(math.Max(float64(colMax[col]), float64(grid[row][col])))
		}
	}
	for _, v := range rowMax {
		xz += v
	}
	for _, v := range colMax {
		yz += v
	}
	return xy + xz + yz
}

```


```python
class Solution:
    def projectionArea(self, grid: List[List[int]]) -> int:
        xy = sum(grid[i][j]>0  for i in range(len(grid)) for j in range(len(grid[0])))
        xz = sum(max(i)  for i in grid)
        yz = sum(max(i)  for i in zip(*grid))

        return xy+xz+yz

```