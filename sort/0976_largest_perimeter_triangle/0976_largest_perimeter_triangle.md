976. Largest Perimeter Triangle


Easy


Given an array A of positive lengths, return the largest perimeter of a triangle with non-zero area, formed from 3 of these lengths.

If it is impossible to form any triangle of non-zero area, return 0.

 

Example 1:
```
Input: [2,1,2]
Output: 5
```


Example 2:
```
Input: [1,2,1]
Output: 0
```

Example 3:

```
Input: [3,2,3,4]
Output: 10
```


Example 4:

```
Input: [3,6,2,3]
Output: 8
```
 

Note:

3 <= A.length <= 10000

1 <= A[i] <= 10^6

## 方法

```go
func largestPerimeter(A []int) int {
    sort.Ints(A)

	for i := len(A) - 1; i >= 0; i-- {
		if i-2 < 0 {
			return 0
		}
		if A[i] < A[i-1]+A[i-2] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
    
	return 0
}
}
```


```python
class Solution(object):
    def largestPerimeter(self, A):
        """
        :type A: List[int]
        :rtype: int
        """
        A.sort()
        for i in reversed(xrange(len(A) - 2)):
            if A[i] + A[i+1] > A[i+2]:
                return A[i] + A[i+1] + A[i+2]
        return 0
```