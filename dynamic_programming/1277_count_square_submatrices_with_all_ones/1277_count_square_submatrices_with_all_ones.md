1277. Count Square Submatrices with All Ones


Medium


Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.

 

Example 1:

```
Input: matrix =
[
  [0,1,1,1],
  [1,1,1,1],
  [0,1,1,1]
]
Output: 15
Explanation: 
There are 10 squares of side 1.
There are 4 squares of side 2.
There is  1 square of side 3.
Total number of squares = 10 + 4 + 1 = 15.
```

Example 2:

```
Input: matrix = 
[
  [1,0,1],
  [1,1,0],
  [1,1,0]
]
Output: 7
Explanation: 
There are 6 squares of side 1.  
There is 1 square of side 2. 
Total number of squares = 6 + 1 = 7.
```

Constraints:

1 <= arr.length <= 300   
1 <= arr[0].length <= 300   
0 <= arr[i][j] <= 1


## 方法


```go
func countSquares(matrix [][]int) int {
	res := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i > 0 && j > 0 && matrix[i][j] > 0 {
				matrix[i][j] = Min(matrix[i-1][j], Min(matrix[i][j-1], matrix[i-1][j-1])) + 1
			}
			res += matrix[i][j]
		}
	}

	return res
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```


```python
class Solution:
    def countSquares(self, matrix: List[List[int]]) -> int:
        m, n = len(matrix), len(matrix[0])
        f = [[0] * n for _ in range(m)]
        ans = 0
        for i in range(m):
            for j in range(n):
                if i == 0 or j == 0:
                    f[i][j] = matrix[i][j]
                elif matrix[i][j] == 0:
                    f[i][j] = 0
                else:
                    f[i][j] = min(f[i][j - 1], f[i - 1][j], f[i - 1][j - 1]) + 1
                ans += f[i][j]
        return ans
```