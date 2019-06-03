766. Toeplitz Matrix

Easy

A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element.

Now given an M x N matrix, return True if and only if the matrix is Toeplitz.
 

Example 1:

```
Input:
matrix = [
  [1,2,3,4],
  [5,1,2,3],
  [9,5,1,2]
]
Output: True
Explanation:
In the above grid, the diagonals are:
"[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]".
In each diagonal all elements are the same, so the answer is True.
```

Example 2:

```
Input:
matrix = [
  [1,2],
  [2,2]
]
Output: False
Explanation:
The diagonal "[1, 2]" has different elements.
```

Note:

matrix will be a 2D array of integers.

matrix will have a number of rows and columns in range [1, 20].

matrix[i][j] will be integers in range [0, 99].


Follow up:

What if the matrix is stored on disk, and the memory is limited such that you can only load at most one row of the matrix into the memory at once?

What if the matrix is so large that you can only load up a partial row into the memory at once?



# 方法
遍历矩阵，分别比较 m[i][j] 和 m[i-1][j-1]，若不相等，则返回 False。



```go
func isToeplitzMatrix(matrix [][]int) bool {
    m, n := len(matrix), len(matrix[0])

	for i := 0; i+1 < m; i++ {
		for j := 0; j+1 < n; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}

	return true
}
```


```python
class Solution(object):
    def isToeplitzMatrix(self, matrix):
        """
        :type matrix: List[List[int]]
        :rtype: bool
        """
        m = len(matrix)
        n = len(matrix[0])
        for i in range(1, m):
          for j in range(1, n):
            if matrix[i][j] != matrix[i - 1][j - 1]: return False
        return True
        
```