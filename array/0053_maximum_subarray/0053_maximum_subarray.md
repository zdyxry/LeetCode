53. Maximum Subarray

Easy


Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:
```
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

# 方法
遍历相加当前值，若当前值与前一个数值之和大于当前值，则当前最大值为当前值与前一个数值相加，否则当前最大值为当前值。
全部数组最大值为当前最大值与数组中最大值取大。

https://www.youtube.com/watch?v=7J5rs56JBs8

```go
package main

import (
	"fmt"
)

func test(nums []int) int {
	f := make([]int, len(nums))
	f[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if f[i-1]+nums[i] < nums[i] {
			f[i] = nums[i]
		} else {
			f[i] = f[i-1] + nums[i]
		}
	}
	res := -1 << 31
	for _, i := range f {
		res = max(res, i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(test(nums))
}

```


```go
package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, sumSoFar := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > sumSoFar+nums[i] {
			sumSoFar = nums[i]
		} else {
			sumSoFar += nums[i]
		}
		if sumSoFar > maxSum {
			maxSum = sumSoFar
		}
	}

	return maxSum
}
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
```

```python
class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        max_nums = max(nums)
        if max_nums < 0:
            return max_nums
        global_max, local_max = 0, 0
        for x in nums:
            local_max = max(0, local_max + x)
            global_max = max(global_max, local_max)
        return global_max
```