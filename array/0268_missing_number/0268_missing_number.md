268. Missing Number

Easy

Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.

Example 1:

```
Input: [3,0,1]
Output: 2
```

Example 2:

```
Input: [9,6,4,2,3,5,7,0,1]
Output: 8
```
Note:
Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?

# 方法

使用异或方式获取不同数字。

```go
func missingNumber(nums []int) int {
    xor := 0

	for i, n := range nums {
		xor ^= i ^ n
	}

	return xor ^ len(nums)
}
```


```python
class Solution(object):
    def missingNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        import operator
        return reduce(operator.xor, nums, \
                      reduce(operator.xor, xrange(len(nums) + 1)))
```