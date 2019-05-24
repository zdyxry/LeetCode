581. Shortest Unsorted Continuous Subarray

Easy

Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order, too.

You need to find the shortest such subarray and output its length.

Example 1:
```
Input: [2, 6, 4, 8, 10, 9, 15]
Output: 5
Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
```

Note:
Then length of the input array is in range [1, 10,000].

The input array may contain duplicates, so ascending order here means <=.

# 方法
遍历数组，维护两个变量，分别保存数组中最大值和最小值，将最大值置为 nums[0]，当数组从前向后遍历过程中，若最大值小于当前值，则数组为递增，符合预期，最大值赋值为当前值，若不小于，则需要调整，记录该数值位置；最小值同理，从后向前遍历，若小于当前值，则不符合预期。



```go
func findUnsortedSubarray(nums []int) int {
    n := len(nums)
	left, right := 0, -1 
	min, max := nums[n-1], nums[0]

	for i := 1; i < n; i++ {
		if max <= nums[i] { 
			max = nums[i]
		} else { 
			right = i
		}

		j := n - i - 1
		if min >= nums[j] {
			min = nums[j]
		} else {
			left = j
		}
	}

	return right - left + 1
}
```

```python
class Solution(object):
    def findUnsortedSubarray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)
        left, right = -1, -2
        min_from_right, max_from_left = nums[-1], nums[0]
        for i in xrange(1, n):
            max_from_left = max(max_from_left, nums[i])
            min_from_right = min(min_from_right, nums[n-1-i])
            if nums[i] < max_from_left: right = i
            if nums[n-1-i] > min_from_right: left = n-1-i
        return right - left + 1
```