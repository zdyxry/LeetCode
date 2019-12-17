378. Kth Smallest Element in a Sorted Matrix


Medium


Given a n x n matrix where each of the rows and columns are sorted in ascending order, find the kth smallest element in the matrix.

Note that it is the kth smallest element in the sorted order, not the kth distinct element.

Example:

```
matrix = [
   [ 1,  5,  9],
   [10, 11, 13],
   [12, 13, 15]
],
k = 8,

return 13.
```

Note:   
You may assume k is always valid, 1 ≤ k ≤ n2.


## 方法

```go
func kthSmallest(mat [][]int, k int) int {
    n := len(mat)

	lo, hi := mat[0][0], mat[n-1][n-1]

	for lo < hi {
		mid := lo + (hi-lo)>>1
		count := 0
		j := n - 1
		for i := 0; i < n; i++ {
			for j >= 0 && mat[i][j] > mid {
				j--
			}
			count += j + 1
		}

		// 移动 lo 或 hi
		if count < k {
			lo = mid + 1
		} else {
			// 没有 -1
			hi = mid
		}
	}

	return lo
}
```

```go
func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	pq := &pq{data: make([]interface{}, k)}
	heap.Init(pq)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if pq.Len() < k {
				heap.Push(pq, matrix[i][j])
			} else if matrix[i][j] < pq.Head().(int) {
				heap.Pop(pq)
				heap.Push(pq, matrix[i][j])
			} else {
				break
			}
		}
	}
	return heap.Pop(pq).(int)
}

type pq struct {
	data []interface{}
	len  int
}

func (p *pq) Len() int {
	return p.len
}

func (p *pq) Less(a, b int) bool {
	return p.data[a].(int) > p.data[b].(int)
}

func (p *pq) Swap(a, b int) {
	p.data[a], p.data[b] = p.data[b], p.data[a]
}

func (p *pq) Push(o interface{}) {
	p.data[p.len] = o
	p.len++
}

func (p *pq) Head() interface{} {
	return p.data[0]
}

func (p *pq) Pop() interface{} {
	p.len--
	return p.data[p.len]
}
```

```python
class Solution(object):
    def kthSmallest(self, matrix, k):
        """
        :type matrix: List[List[int]]
        :type k: int
        :rtype: int
        """
        left = matrix[0][0]
        right = matrix[-1][-1] + 1
        while left < right:
            mid = (left+right)/2 
            i = len(matrix)-1 
            j = 0 
            count = 0 
            while i >= 0 and  j < len(matrix[0]):
                if matrix[i][j] <= mid:
                    j += 1
                    count += (i+1)
                else:
                    i -= 1
            
            if count < k:
                left = mid + 1 
            else:
                right = mid 
            
        return left 
```