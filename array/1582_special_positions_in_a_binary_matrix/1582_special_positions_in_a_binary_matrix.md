1582. Special Positions in a Binary Matrix


Easy


Given a rows x cols matrix mat, where mat[i][j] is either 0 or 1, return the number of special positions in mat.

A position (i,j) is called special if mat[i][j] == 1 and all other elements in row i and column j are 0 (rows and columns are 0-indexed).

 

Example 1:

```
Input: mat = [[1,0,0],
              [0,0,1],
              [1,0,0]]
Output: 1
Explanation: (1,2) is a special position because mat[1][2] == 1 and all other elements in row 1 and column 2 are 0.
```

Example 2:

```
Input: mat = [[1,0,0],
              [0,1,0],
              [0,0,1]]
Output: 3
Explanation: (0,0), (1,1) and (2,2) are special positions. 
```

Example 3:

```
Input: mat = [[0,0,0,1],
              [1,0,0,0],
              [0,1,1,0],
              [0,0,0,0]]
Output: 2
```

Example 4:

```
Input: mat = [[0,0,0,0,0],
              [1,0,0,0,0],
              [0,1,0,0,0],
              [0,0,1,0,0],
              [0,0,0,1,1]]
Output: 3
```
 

Constraints:

rows == mat.length  
cols == mat[i].length  
1 <= rows, cols <= 100  
mat[i][j] is 0 or 1.


## 方法


```go
func numSpecial(mat [][]int) int {
	m, n, ans := len(mat), len(mat[0]), 0
	row , col := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++{
			if mat[i][j] == 1{
				row[i]++
				col[j]++
			}
		}
	}

	pool := []int{}
	for j := 0; j < n; j++ {
		if col[j] == 1{
			pool = append(pool, j)
		}
	}

	for i := 0; i < m; i++ {
		if row[i] != 1 {
			continue
		}
		for j := range pool {
			if mat[i][pool[j]] == 1 {
				ans++
				break
			}
		}
	}
	return ans
}
```

```python
class Solution:
    def numSpecial(self, mat: List[List[int]]) -> int:
        m, n, ans = len(mat), len(mat[0]), 0
        row, col = [0] * m, [0] * n
        for i in range(m):
            for j in range(n):
                if mat[i][j]:
                    row[i] += 1
                    col[j] += 1
        pool = [j for j in range(n) if col[j] == 1]
        for i in range(m):
            if row[i] != 1:
                continue
            for j in pool:
                if mat[i][j]:
                    ans += 1
                    break
        return ans
```