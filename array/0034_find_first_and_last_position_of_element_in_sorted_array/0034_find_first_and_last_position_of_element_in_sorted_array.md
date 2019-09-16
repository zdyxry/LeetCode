34. Find First and Last Position of Element in Sorted Array


Medium


Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:
```
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```


Example 2:
```
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```


## 方法


```go
func searchRange(nums []int, target int) []int {
    index := search(nums, target)
	if index == -1 {
		return []int{-1, -1}
	}

	// 利用二分法，查找第一个target
	first := index
	for {
		f := search(nums[:first], target)
		if f == -1 {
			break
		}
		first = f
	}

	// 利用二分法，查找最后一个target
	last := index
	for {
		l := search(nums[last+1:], target)
		if l == -1 {
			break
		}
		// 注意这里与查找first的不同
		last += l + 1
	}

	return []int{first, last}
}

// 二分查找法
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	var median int

	for low <= high {
		median = (low + high) / 2

		switch {
		case nums[median] < target:
			low = median + 1
		case nums[median] > target:
			high = median - 1
		default:
			return median
		}
	}
	return -1
}
```


```python
class Solution(object):
    def searchRange(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        def get_start_index(nums, target):
            start, end = 0, len(nums) - 1
            while start <= end:
                mid = start + (end - start) // 2
                if nums[mid] == target:
                    if mid == 0 or nums[mid-1] != target:
                        return mid
                if nums[mid] < target:
                    start = mid + 1
                else:
                    end = mid - 1
            return -1
        
        def get_end_index(nums, target):
            start, end = 0, len(nums) - 1
            while start <= end:
                mid = start + (end - start)//2
                if nums[mid] == target: 
                    if mid == len(nums)-1 or nums[mid + 1] != target:
                        return mid
                if nums[mid] <= target:
                    start = mid + 1
                else:
                    end = mid - 1
            return -1
        
        return [get_start_index(nums, target), get_end_index(nums, target)]
```