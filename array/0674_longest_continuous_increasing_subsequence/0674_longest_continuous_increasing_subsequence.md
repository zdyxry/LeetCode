674. Longest Continuous Increasing Subsequence

Easy

Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).

Example 1:
```
Input: [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3. 
Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4. 
```
Example 2:
```
Input: [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2], its length is 1. 
```

Note: Length of the array will not exceed 10,000.


# 方法
遍历数组，当当前数值大于前一个数值时，取间隔最大值并返回。


```go
func findLengthOfLCIS(nums []int) int {
	res, cur := 0, 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] < nums[i] {
			cur++
		} else {
			res = int(math.Max(float64(res), float64(cur)))
			cur = 1
		}
	}
	return int(math.Max(float64(res), float64(cur)))
}
```


```python

class Solution(object):
    def findLengthOfLCIS(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        result, count = 0, 0
        for i in xrange(len(nums)):
            if i == 0 or nums[i-1] < nums[i]:
                count += 1
                result = max(result, count)
            else:
                count = 1
        return result
```