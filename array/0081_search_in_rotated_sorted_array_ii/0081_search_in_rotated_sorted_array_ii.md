81. Search in Rotated Sorted Array II


Medium


Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).

You are given a target value to search. If found in the array return true, otherwise return false.

Example 1:

```
Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true
```

Example 2:
```
Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false
```

Follow up:

This is a follow up problem to Search in Rotated Sorted Array, where nums may contain duplicates.

Would this affect the run-time complexity? How and why?

## 方法

```go
func search(nums []int, target int) bool {
    length := len(nums)
	if length == 0 {
		return false
	}
	k := 1
	for k < len(nums) && nums[k-1] <= nums[k] {
		k++
	}

	i, j := 0, length-1
	for i <= j {
		m := (i + j) / 2
		med := (m + k) % length

		switch {
		case nums[med] < target:
			i = m + 1
		case target < nums[med]:
			j = m - 1
		default:
			return true
		}
	}

	return false
}
```



```python
class Solution(object):
    def search(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: bool
        """
        left, right = 0, len(nums) - 1

        while left <= right:
            mid = left + (right - left) / 2

            if nums[mid] == target:
                return True
            elif nums[mid] == nums[left]:
                left += 1
            elif (nums[mid] > nums[left] and nums[left] <= target < nums[mid]) or \
                 (nums[mid] < nums[left] and not (nums[mid] < target <= nums[right])):
                right = mid - 1
            else:
                left = mid + 1

        return False
```