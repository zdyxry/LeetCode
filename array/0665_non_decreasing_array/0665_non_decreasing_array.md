665. Non-decreasing Array

Easy

Given an array with n integers, your task is to check if it could become non-decreasing by modifying at most 1 element.

We define an array is non-decreasing if array[i] <= array[i + 1] holds for every i (1 <= i < n).

Example 1:
```
Input: [4,2,3]
Output: True
Explanation: You could modify the first 4 to 1 to get a non-decreasing array.
```
Example 2:
```
Input: [4,2,1]
Output: False
Explanation: You can't get a non-decreasing array by modify at most one element.
```

Note: The n belongs to [1, 10,000].

# 方法
遍历一次数组，检查当前值 nums[i] 和 nums[i-1] nums[i-2] 关系，决定调整哪个数值来保证数组非递减。


```go
func checkPossibility(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 || len(nums) == 2 {
		return true
	}

	found := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if found {
				return false
			}
			if i-2 >= 0 && nums[i] < nums[i-2] {
				nums[i] = nums[i-1]
			}
			found = true
		}
	}

	return true
}
```


```python
class Solution(object):
    def checkPossibility(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        cnt = 0
        for i in range(1, len(nums)):
            if nums[i - 1] > nums[i]:
                cnt += 1
                if i < 2 or nums[i - 2] <= nums[i]:
                    nums[i - 1] = nums[i] # modify nums[i-1]
                else:
                    nums[i] = nums[i - 1] # modify nums[i]
        return cnt <= 1
```