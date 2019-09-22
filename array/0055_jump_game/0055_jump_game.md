55. Jump Game


Medium


Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:
```
Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```


Example 2:
```
Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.
```

## 方法

```go
func canJump(nums []int) bool {
    for i := len(nums) - 2; i >= 0; i-- {
		// 找到数值为 0 的元素
		if nums[i] != 0 {
			continue
		}

		j := i - 1
		for ; j >= 0; j-- {
			if i-j < nums[j] {
				// 在 j 号位置上，可以跨过 0 元素
				i = j
				break
			}
		}

		if j == -1 {
			// 在 0 元素之前，没有位置可以跨过 0
			return false
		}
	}

	return true
}
```


```python
class Solution(object):
    def canJump(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        reachable = 0
        for i, length in enumerate(nums):
            if i > reachable:
                break
            reachable = max(reachable, i + length)
        return reachable >= len(nums) - 1
```