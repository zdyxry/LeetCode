498. Diagonal Traverse
Medium


Given a matrix of M x N elements (M rows, N columns), return all elements of the matrix in diagonal order as shown in the below image.

 

Example:

```
Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]

Output:  [1,2,4,7,5,3,6,8,9]
```

Explanation:

![1](0498-1.png)


## 方法

```go
func findDiagonalOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	n, m := len(matrix), len(matrix[0])
	res := make([]int, n*m)
	var x, y int
	for i := 0; i < len(res); i++ {
		res[i] = matrix[x][y]
		if (x+y)%2 == 0 {
			if y == m-1 {
				x++
			} else if x == 0 {
				y++
			} else {
				y++
				x--
			}
		} else {
			if x == n-1 {
				y++
			} else if y == 0 {
				x++
			} else {
				x++
				y--
			}
		}
	}
	return res
}
```


```python
import collections

class Solution:
    def findDiagonalOrder(self, matrix: List[List[int]]) -> List[int]:
        map = {}
        res = []
        for r, _ in enumerate(matrix):
            for c, _ in enumerate(matrix[r]):
                if not map.get(r+c):
                    map[r+c] = []
                map[r+c].append(matrix[r][c])
        od = collections.OrderedDict(sorted(map.items()))
        flag = False
        for k, v in od.items():
            if flag:
                res.extend(v)
                flag = False
            else:
                res.extend(reversed(v))
                flag = True
        return res
```