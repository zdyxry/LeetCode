118. Pascal's Triangle

Easy

Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

```
Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
```

# 方法
维护两个索引，一个为当前行，一个为当前列，通过上一行的数值算出当前行数值，进行数组拼接。

```go
func generate(numRows int) [][]int {
	res := [][]int{}
	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}

	for i := 1; i < numRows; i++ {
		res = append(res, genNext(res[i-1]))
	}

	return res
}

func genNext(p []int) []int {
	res := make([]int, 1, len(p)+1)
	res = append(res, p...)

	for i := 0; i < len(res)-1; i++ {
		res[i] += res[i+1]
	}

	return res
}
```


```python
class Solution(object):
    def generate(self, numRows):
        """
        :type numRows: int
        :rtype: List[List[int]]
        """
        result = []
        for i in range(numRows):
            result.append([])
            for j in range(0, i+1):
                if j in (0,i):
                    result[i].append(1)
                else:
                    result[i].append(result[i-1][j-1] + result[i-1][j])
                    
        return result
```