448. Find All Numbers Disappeared in an Array

Easy


Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:

```
Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
```

# 方法
遍历数组，将当前数值绝对值索引处数值转换为负数，再次遍历数组，若数值为整数，则确实数字为该整数索引 + 1。




```go
func findDisappearedNumbers(nums []int) []int {
    for i := 0; i < len(nums); i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	res := make([]int, 0, len(nums))

	for i, n := range nums {
		if n != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}
```


```python
class Solution(object):
    def findDisappearedNumbers(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        res=[1]*len(nums)
        for i in nums:
            res[i-1]=None
        return [i+1 for i in range(len(nums)) if res[i]]
```


```python
class Solution(object):
    def findDisappearedNumbers(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        for i in xrange(len(nums)):
            if nums[abs(nums[i]) - 1] > 0:
                nums[abs(nums[i]) - 1] *= -1

        result = []
        for i in xrange(len(nums)):
            if nums[i] > 0:
                result.append(i+1)
            else:
                nums[i] *= -1
        return result
```