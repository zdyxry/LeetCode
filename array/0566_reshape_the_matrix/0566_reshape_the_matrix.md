566. Reshape the Matrix
 
Easy

In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.

You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.

The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.

If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

Example 1:
```
Input: 
nums = 
[[1,2],
 [3,4]]
r = 1, c = 4
Output: 
[[1,2,3,4]]
Explanation:
The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix, fill it row by row by using the previous list.
```
Example 2:
```
Input: 
nums = 
[[1,2],
 [3,4]]
r = 2, c = 4
Output: 
[[1,2],
 [3,4]]
Explanation:
There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So output the original matrix.
```

Note:
The height and width of the given matrix is in range [1, 100].

The given r and c are all positive.

# 方法
遍历二维数组，通过对列整除和取余确定下标。





```go
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || len(nums[0]) == 0 || len(nums)*len(nums[0]) != r*c || len(nums) == r && len(nums[0]) == c {
		return nums
	}

	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}
	m, n := len(nums), len(nums[0])
	for i := 0; i < m*n; i++ {
		org_y := i / n
		org_x := i % n
		new_y := i / c
		new_x := i % c
		res[new_y][new_x] = nums[org_y][org_x]
	}

	return res
}
```



```go
func matrixReshape(nums [][]int, r int, c int) [][]int {
    if len(nums) == 0 || len(nums[0]) == 0 || len(nums)*len(nums[0]) != r*c || len(nums) == r && len(nums[0]) == c {
		return nums
	}

	res := make([][]int, r)
	count, col := 0, len(nums[0])
	for i := range res {
		res[i] = make([]int, c)

		for j := range res[i] {
			res[i][j] = nums[count/col][count%col]
			count++
		}
	}

	return res
}
```


```python
class Solution(object):
    def matrixReshape(self, nums, r, c):
        """
        :type nums: List[List[int]]
        :type r: int
        :type c: int
        :rtype: List[List[int]]
        """
        if not nums or \
           r*c != len(nums) * len(nums[0]):
            return nums

        result = [[0 for _ in xrange(c)] for _ in xrange(r)]
        count = 0
        for i in xrange(len(nums)):
            for j in xrange(len(nums[0])):
                result[count/c][count%c] = nums[i][j]
                count += 1
        return result
            
```