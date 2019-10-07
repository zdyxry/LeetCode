120. Triangle


Medium


Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

```
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
```

Note:

Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.


## 方法

```go
func minimumTotal(triangle [][]int) int {
    n := len(triangle)
	if n == 0 {
		return 0
	}

	// 从上往下，依次修改triangle
	// 使得 triangle[i][j] 表示到达 (i,j) 点的最小值
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			switch {
			case j == 0:
				triangle[i][0] += triangle[i-1][0]
			case j == i:
				triangle[i][i] += triangle[i-1][i-1]
			case triangle[i-1][j-1] < triangle[i-1][j]:
				triangle[i][j] += triangle[i-1][j-1]
			default:
				triangle[i][j] += triangle[i-1][j]
			}
		}
	}

	// 从最后一行，获取最小值
	min := triangle[n-1][0]
	for j := 1; j < n; j++ {
		if min > triangle[n-1][j] {
			min = triangle[n-1][j]
		}
	}

	return min
}
```


```python
class Solution(object):
    def minimumTotal(self, triangle):
        """
        :type triangle: List[List[int]]
        :rtype: int
        """
        n = len(triangle)
        if n==0: return None
        if n==1: return triangle[0][0]
        for i in range(1, n):
            triangle[i][0] += triangle[i-1][0]
            triangle[i][-1] += triangle[i-1][-1]
            for j in range(1, i):
                triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
        return min(triangle[-1])
```