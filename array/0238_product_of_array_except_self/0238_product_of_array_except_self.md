238. Product of Array Except Self


Medium


Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].

Example:

```
Input:  [1,2,3,4]
Output: [24,12,8,6]
```

Note: Please solve it without division and in O(n).

Follow up:
Could you solve it with constant space complexity? (The output array does not count as extra space for the purpose of space complexity analysis.)

## 方法

```go
func productExceptSelf(nums []int) []int {
    size := len(nums)

	res := make([]int, size)

	left := 1 // product all left of nums[i]
	for i := 0; i < size; i++ {
		res[i] = left
		left *= nums[i]
	}

	right := 1 // product all right of nums[i]
	for i := size - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}
```


```python
class Solution(object):
    def productExceptSelf(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        if not nums:
            return []

        left_product = [1 for _ in xrange(len(nums))]
        for i in xrange(1, len(nums)):
            left_product[i] = left_product[i - 1] * nums[i - 1]

        right_product = 1
        for i in xrange(len(nums) - 2, -1, -1):
            right_product *= nums[i + 1]
            left_product[i] = left_product[i] * right_product

        return left_product
```