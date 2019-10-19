260. Single Number III


Medium


Given an array of numbers nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.

Example:

```
Input:  [1,2,1,3,2,5]
Output: [3,5]
```

Note:

The order of the result is not important. So in the above example, [5, 3] is also correct.

Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?

## 方法

```go
func singleNumber(nums []int) []int {
    var xor int
	for _, num := range nums {
		xor ^= num
	}

	lowest := xor & -xor

	var a, b int
	for _, num := range nums {
		if num&lowest == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
```



```python
class Solution(object):
    def singleNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        x_xor_y = 0
        for i in nums:
            x_xor_y ^= i

        bit = x_xor_y & ~(x_xor_y - 1)

        x = 0
        for i in nums:
            if i & bit:
                x ^= i

        return [x, x ^ x_xor_y]
```