1380. Lucky Numbers in a Matrix


Easy


Given a m * n matrix of distinct numbers, return all lucky numbers in the matrix in any order.

A lucky number is an element of the matrix such that it is the minimum element in its row and maximum in its column.

 

Example 1:

```
Input: matrix = [[3,7,8],[9,11,13],[15,16,17]]
Output: [15]
Explanation: 15 is the only lucky number since it is the minimum in its row and the maximum in its column
```

Example 2:

```
Input: matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
Output: [12]
Explanation: 12 is the only lucky number since it is the minimum in its row and the maximum in its column.
```

Example 3:

```
Input: matrix = [[7,8],[1,2]]
Output: [7]
```
 

Constraints:

m == mat.length  
n == mat[i].length  
1 <= n, m <= 50  
1 <= matrix[i][j] <= 10^5.  
All elements in the matrix are distinct.  

## 方法


```go
func luckyNumbers (matrix [][]int) []int {
    rows, cols := len(matrix), len(matrix[0])
    res := make([]int, 0)
    for i := 0; i < rows; i++{
        //行最小
        left, right := 0, cols-1
        for left < right{
            if matrix[i][left] > matrix[i][right] {
                left++
            } else {
                right--
            }
        }
        //列最大
        up, down := 0, rows-1
        for up < down{
            if matrix[up][left] > matrix[down][left]{
                down--
            } else {
                up++
            }
        }
        //是不是同一行
        if i == up {
            res = append(res, matrix[up][left])
        }
    }
    return res

}
```



```python
class Solution(object):
    def luckyNumbers (self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: List[int]
        """
        mins = {min(rows) for rows in matrix}
        maxes = {max(columns) for columns in zip(*matrix)}
        return list(mins & maxes)
```
