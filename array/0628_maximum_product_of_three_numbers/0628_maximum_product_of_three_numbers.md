628. Maximum Product of Three Numbers

Easy

Given an integer array, find three numbers whose product is maximum and output the maximum product.

Example 1:

```
Input: [1,2,3]
Output: 6
```
 

Example 2:

```
Input: [1,2,3,4]
Output: 24
 
```

Note:

The length of the given array will be in range [3,104] and all elements are in the range [-1000, 1000].

Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.


# 方法
遍历数组，取数组中最大的 3 个元素。
考虑数字为负数情况：

max1 | max2 | max3 | min1 | min2  | result
--- | --- | --- | --- | --- | ---
正数 | 正数| 正数| 正数 | 正数 | max1 * max2 * max3
负数 | 负数| 负数| 负数| 负数| max1 * max2 * max3
正数 | 正数| 负数| 负数| 负数| max1 * max2 * max3
正数 | 负数 | 负数| 负数| 负数 | max1 * min1 * min2



```go

func maximumProduct(nums []int) int {

	max := -1001
	max1 := -1001
	max2 := -1001
	min1 := 1001
	min2 := 1001

	for _, n := range nums {
		switch {
		case n > max:
			max2, max1, max = max1, max, n
		case n > max1:
			max2, max1 = max1, n
		case n > max2:
			max2 = n
		}

		switch {
		case n < min1:
			min2, min1 = min1, n
		case n < min2:
			min2 = n
		}
	}

	return bigger(max1*max2, min1*min2) * max
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```

```python
class Solution(object):
    def maximumProduct(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        min1, min2 = float('inf'), float('inf')
        max1, max2, max3 = float('-inf'), float('-inf'), float('-inf')
        
        for num in nums:
            if num > max1:
                max3 = max2
                max2 = max1
                max1 = num
            elif num > max2:
                max3 = max2
                max2 = num
            elif num > max3:
                max3 = num
            
            if num < min1:
                min2 = min1
                min1 = num
            elif num < min2:
                min2 = num
            
        return max(max3*max2, min1*min2) * max1
```