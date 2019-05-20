485. Max Consecutive Ones

Easy


Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:

```
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.
```


Note:

The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000

# 方法
遍历数组，当数字为 1 时计数加一，取最大计数。


```go
func findMaxConsecutiveOnes(nums []int) int {
    max, count := 0, 0
	for _, v := range nums {
		switch v {
		case 1:
			count++
			if count > max {
				max = count
			}
		case 0:
			count = 0
		}
	}

	return max
}
```


```python
class Solution(object):
    def findMaxConsecutiveOnes(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        cnt = 0
        ans = 0
        for num in nums:
            if num == 1:
                cnt += 1
                ans = max(ans, cnt)
            else:
                cnt = 0
        return ans
```