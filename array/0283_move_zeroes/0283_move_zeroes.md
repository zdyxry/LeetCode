283. Move Zeroes

Easy

Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

```
Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
```

Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.

# 方法
同时保存两个指针，其中 1 个为非 0 值索引，另一个为正常遍历索引，当数值不为 0 时，则将其交换至非 0 索引所在位置。



```go
func moveZeroes(nums []int)  {
    pos := 0
    for idx, _ := range nums {
        if nums[idx] != 0 {
            nums[idx], nums[pos] = nums[pos], nums[idx]
            pos += 1
        }
    }
}
```


```python
class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        pos = 0
        for i in range(len(nums)):
            if nums[i]:
                nums[i], nums[pos] = nums[pos], nums[i]
                pos += 1
```