74. Search a 2D Matrix


Medium


Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

Integers in each row are sorted from left to right.
The first integer of each row is greater than the last integer of the previous row.
Example 1:

```
Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
Output: true
```


Example 2:

```
Input:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
Output: false
```


## 方法

```go
func searchMatrix(mat [][]int, target int) bool {
	m := len(mat)
	if m == 0 {
		return false
	}

	n := len(mat[0])
	if n == 0 {
		return false
	}

	if target < mat[0][0] || mat[m-1][n-1] < target {
		return false
	}

	// 寻找行
	r := 0
	for r < m && mat[r][0] <= target {
		r++
	}
	r--

	// 二分法寻找 target
	i, j := 0, n-1
	for i <= j {
		med := (i + j) / 2
		switch {
		case mat[r][med] < target:
			i = med + 1
		case target < mat[r][med]:
			j = med - 1
		default:
			return true
		}
	}

	return mat[r][j] == target
}
```


```python
class Solution(object):
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        if not matrix or target is None:
            return False

        rows, cols = len(matrix), len(matrix[0])
        low, high = 0, rows * cols - 1
        
        while low <= high:
            mid = (low + high) / 2
            num = matrix[mid / cols][mid % cols]

            if num == target:
                return True
            elif num < target:
                low = mid + 1
            else:
                high = mid - 1
        
        return False
```