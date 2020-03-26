1329. Sort the Matrix Diagonally


Medium


Given a m * n matrix mat of integers, sort it diagonally in ascending order from the top-left to the bottom-right then return the sorted array.

 

Example 1:

![1329](1329.png)

```
Input: mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
Output: [[1,1,1,1],[1,2,2,2],[1,2,3,3]]
```

Constraints:

m == mat.length  
n == mat[i].length  
1 <= m, n <= 100  
1 <= mat[i][j] <= 100


## 方法

```go
import sort2 "sort"

func diagonalSort(mat [][]int) [][]int {
	var (
		i    int
		lenX = len(mat)
	)
	if lenX == 0 {
		return mat
	}

	var (
		lenY = len(mat[0])
	)

	for i = 0; i < lenX; i++ {
		sort(mat, lenX-1-i, 0, lenX, lenY)
	}
	for i = 1; i < lenY; i++ {
		sort(mat, 0, i, lenX, lenY)
	}
	return mat
}

func sort(mat [][]int, x, y, lenX, lenY int) {
	var (
		tmp        []int
		tmpX, tmpY = x, y
	)
	for tmpX < lenX && tmpY < lenY {
		tmp = append(tmp, mat[tmpX][tmpY])
		tmpX++
		tmpY++
	}
	tmpX = 0
	sort2.Ints(tmp)
	for x < lenX && y < lenY {
		mat[x][y] = tmp[tmpX]
		tmpX++
		x++
		y++
	}
}

```



```python
class Solution(object):
    def diagonalSort(self, mat):
        """
        :type mat: List[List[int]]
        :rtype: List[List[int]]
        """
        diag = defaultdict(list)
        N,M = len(mat), len(mat[0])
        
        for i in range(N):
            for j in range(M):
                diag[i-j].append(mat[i][j])
        
        for k in diag.keys():
            diag[k].sort(reverse=True)
        
        for i in range(N):
            for j in range(M):
                mat[i][j] = diag[i-j].pop()
        
        return mat
```