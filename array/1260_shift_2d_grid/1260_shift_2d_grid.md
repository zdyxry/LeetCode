1260. Shift 2D Grid


Easy


Given a 2D grid of size m x n and an integer k. You need to shift the grid k times.

In one shift operation:

```
Element at grid[i][j] moves to grid[i][j + 1].
Element at grid[i][n - 1] moves to grid[i + 1][0].
Element at grid[m - 1][n - 1] moves to grid[0][0].
Return the 2D grid after applying shift operation k times.
```
 

Example 1:


```
Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 1
Output: [[9,1,2],[3,4,5],[6,7,8]]
```

Example 2:


```
Input: grid = [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]], k = 4
Output: [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
```

Example 3:

```
Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 9
Output: [[1,2,3],[4,5,6],[7,8,9]]
```

Constraints:

m == grid.length  
n == grid[i].length  
1 <= m <= 50  
1 <= n <= 50  
-1000 <= grid[i][j] <= 1000  
0 <= k <= 100  

## 方法

```go
func shiftGrid(grid [][]int, k int) [][]int {
    res := make([]int,0,10)
    lenth := len(grid)
    w := len(grid[0])
    for i:=0;i<lenth;i++{
        res = append(res,grid[i]...)
    }
    k = k % (lenth*w)
    res = append(res[lenth*w-k:],res[:lenth*w-k]...)
    tmp := make([][]int,0,10)
    for i:=0;i<lenth;i++{
        tmp = append(tmp,res[i*w:(i+1)*w])
    }
    return tmp
}
```


```python
class Solution:
    def shiftGrid(self, grid: List[List[int]], k: int) -> List[List[int]]:
        num_rows, num_cols = len(grid), len(grid[0])

        for _ in range(k):

            previous = grid[-1][-1]
            for row in range(num_rows):
                for col in range(num_cols):
                    temp = grid[row][col]
                    grid[row][col] = previous
                    previous = temp
        return grid
```