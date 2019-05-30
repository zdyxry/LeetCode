697. Degree of an Array

Easy

Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of any one of its elements.

Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.

Example 1:
```j
Input: [1, 2, 2, 3, 1]
Output: 2
Explanation: 
The input array has a degree of 2 because both elements 1 and 2 appear twice.
Of the subarrays that have the same degree:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
The shortest length is 2. So return 2.
```
Example 2:
```
Input: [1,2,2,3,1,4,2]
Output: 6
```

Note:

nums.length will be between 1 and 50,000.

nums[i] will be an integer between 0 and 49,999.

# 方法

遍历数组，分别保存数字出现的第一次位置和数字出现的总次数，当数字出现次数与最多次数想等时，比较索引位置最小值并返回。


```go
func findShortestSubArray(nums []int) int {
    size := len(nums)
	if size < 2 {
		return size
	}

	first := make(map[int]int, size)
	count := make(map[int]int, size)
	maxCount := 1
	minLen := size
	for i, n := range nums {
		count[n]++
		if count[n] == 1 {
			first[n] = i
		} else {
			l := i - first[n] + 1
			if maxCount < count[n] ||
				(maxCount == count[n] && minLen > l) {
				maxCount = count[n]
				minLen = l
			}
		}
	}

	if len(count) == size {
		return 1
	}
	return minLen
}
```



```python
class Solution(object):
    def findShortestSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        first, counter, res, degree = {}, {}, 0, 0
        for i, v in enumerate(nums):
            first.setdefault(v, i)
            counter[v] = counter.get(v, 0) + 1
            if counter[v] > degree:
                degree = counter[v]
                res = i - first[v] + 1
            elif counter[v] == degree:
                res = min(res, i - first[v] + 1)
        return res

```