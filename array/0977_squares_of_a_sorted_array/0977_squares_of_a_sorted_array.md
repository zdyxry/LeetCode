977. Squares of a Sorted Array

Easy

Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.

 

Example 1:

```
Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
```
Example 2:

```
Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]
```

Note:

1 <= A.length <= 10000

-10000 <= A[i] <= 10000

A is sorted in non-decreasing order.


# 方法
双指针遍历数组，当最大值与最小值和为负数，则res 最大值为最小值的平方。

```go
func sortedSquares(A []int) []int {
	size := len(A)
	res := make([]int, size)
	for l, r, i := 0, size-1, size-1; l <= r; i-- {
		if A[l]+A[r] < 0 {
			res[i] = A[l] * A[l]
			l++
		} else {
			res[i] = A[r] * A[r]
			r--
		}
	}
	return res
}
```


```python
class Solution(object):
    def sortedSquares(self, A):
        """
        :type A: List[int]
        :rtype: List[int]
        """
        size = len(A)
        res = [0] * size

        l, r, i = 0, size - 1, size - 1
        while l <= r:
            if A[l] + A[r] < 0:
                res[i] = A[l] * A[l]
                l += 1
            else:
                res[i] = A[r] * A[r]
                r -= 1
            i -= 1
        return res
```