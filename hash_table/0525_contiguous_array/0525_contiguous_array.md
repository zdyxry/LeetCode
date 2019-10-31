525. Contiguous Array


Medium


Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.

Example 1:
```
Input: [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
```

Example 2:
```
Input: [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
Note: The length of the given binary array will not exceed 50,000.
```

## 方法


```go
func findMaxLength(nums []int) int {
    delta := make(map[int]int, len(nums)/2)
	delta[0] = -1

	ones, zeros, res := 0, 0, 0

	for i := range nums {
		if nums[i] == 1 {
			ones++
		} else {
			zeros++
		}

		if j, ok := delta[ones-zeros]; ok {
			res = max(res, i-j)
		} else {
			delta[ones-zeros] = i
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```



```python
class Solution(object):
    def findMaxLength(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        result, count = 0, 0
        lookup = {0: -1}
        for i, num in enumerate(nums):
            count += 1 if num == 1 else -1
            if count in lookup:
                result = max(result, i - lookup[count])
            else:
                lookup[count] = i

        return result
```