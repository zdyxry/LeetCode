169. Majority Element

Easy

Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

```
Input: [3,2,3]
Output: 3
```

Example 2:

```
Input: [2,2,1,1,1,2,2]
Output: 2
```

# 方法
1. map 用数字作为 key，次数作为 value，当次数大于 len(nums) / 2，则返回该数
2. 当做多次数的数大于 len(nums)/2，可以转换为每个位上的大多数均为最终值，比如最低位，1 出现次数大于 0 出现次数，则最终值的最低位为 1.
3. 对数组进行排序，则中位数为最终值


```go
func majorityElement(nums []int) int {
	x, t := nums[0], 1

	for i := 1; i < len(nums); i++ {
		switch {
		case x == nums[i]:
			t++
		case t > 0:
			t--
		default:
			x = nums[i]
			t = 1
		}
	}

	return x
}
```


```python
class Solution(object):
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        idx, cnt = 0, 1

        for i in xrange(1, len(nums)):
            if nums[idx] == nums[i]:
                cnt += 1
            else:
                cnt -= 1
                if cnt == 0:
                    idx = i
                    cnt = 1
        return nums[idx]
```