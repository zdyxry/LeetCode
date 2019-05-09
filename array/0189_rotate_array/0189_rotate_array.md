189. Rotate Array

Easy

Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:

```
Input: [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
```

Example 2:
```
Input: [-1,-100,3,99] and k = 2
Output: [3,99,-1,-100]
Explanation: 
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
```

Note:

Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?

# 方法
通过反转数组的方式实现比较好理解，且题目要求空间 O(1)。改变索引方式效率没有大幅提高。

```go
func rotate(nums []int, k int) {
	n := len(nums)

	if k > n {
		k %= n
	}
	if k == 0 || k == n {
		return
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
```


```python
class Solution(object):
    def rotate(self, nums, k):
        k %= len(nums)
        self.reverse(nums, 0, len(nums))
        self.reverse(nums, 0, k)
        self.reverse(nums, k, len(nums))

    def reverse(self, nums, start, end):
        while start < end:
            nums[start], nums[end - 1] = nums[end - 1], nums[start]
            start += 1
            end -= 1
```

```python
class Solution(object):
    def rotate(self, nums, k):
        n = len(nums)
        idx = 0
        distance = 0
        cur = nums[0]
        for x in range(n):
            idx = (idx + k) % n
            nums[idx], cur = cur, nums[idx]
            
            distance = (distance + k) % n
            if distance == 0:
                idx = (idx + 1) % n
                cur = nums[idx]
```