219. Contains Duplicate II

Easy

Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

Example 1:

```
Input: nums = [1,2,3,1], k = 3
Output: true
```

Example 2:

```
Input: nums = [1,0,1,1], k = 1
Output: true
```

Example 3:

```
Input: nums = [1,2,3,1,2,3], k = 2
Output: false
```

# 方法
通过字典记录最近的一次出现索引。


```go
func containsNearbyDuplicate(nums []int, k int) bool {
    if k <= 0 {
		return false
	}
	
	// 利用 m 记录 a[i]中的值，每次出现的(索引号+1)
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if m[n] != 0 && i-(m[n]-1) <= k {
			return true
		}
		m[n] = i + 1
	}

	return false
}
```


```python
class Solution(object):
    def containsNearbyDuplicate(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: bool
        """
        lookup = {}
        for i, num in enumerate(nums):
            if num not in lookup:
                lookup[num] = i
            else:
                if i - lookup[num] <= k:
                    return True
                lookup[num] = i
        return False
```