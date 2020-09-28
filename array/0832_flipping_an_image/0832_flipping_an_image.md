832. Flipping an Image


Easy


Given a binary matrix A, we want to flip the image horizontally, then invert it, and return the resulting image.

To flip an image horizontally means that each row of the image is reversed.  For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].

To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].

Example 1:

```
Input: [[1,1,0],[1,0,1],[0,0,0]]
Output: [[1,0,0],[0,1,0],[1,1,1]]
Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
```

Example 2:

```
Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
```

Notes:

1 <= A.length = A[0].length <= 20  
0 <= A[i][j] <= 1


## 方法


```go
func flipAndInvertImage(A [][]int) [][]int {
    for i, row := range A {
		for j := range row {
			if j > (len(row)-1)/2 {
				break
			}

			tmp := A[i][j]
			A[i][j] = A[i][len(row)-1-j]
			A[i][len(row)-1-j] = tmp

			A[i][j] = A[i][j] ^ 1
			if len(row)-1-j != j {
				A[i][len(row)-1-j] = A[i][len(row)-1-j] ^ 1
			}

		}
	}
	return A
}
```


```python
class Solution:
    def flipAndInvertImage(self, A: List[List[int]]) -> List[List[int]]:
        m, n = len(A), len(A[0])
        for i in range(m):
            j, k = 0, n - 1
            # 行翻转
            while j < k:
                A[i][j], A[i][k] = A[i][k], A[i][j]
                j += 1
                k -= 1

        # 0, 1互换
        for i in range(m):
            for j in range(n):
                A[i][j] = 1 if A[i][j] == 0 else 0

        return A
```