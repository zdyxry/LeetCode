209. Minimum Size Subarray Sum


Medium


Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

Example: 

```
Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.
```

Follow up:

If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n). 

## 方法

```go
func minSubArrayLen(s int, a []int) int {
    n := len(a)
	res, i, j, sum := n+1, 0, 0, 0

	for j < n {
		sum += a[j]
		j++

		for sum >= s {
			sum -= a[i]
			i++
			if res > j-i+1 {
				res = j - i + 1
			}
		}
	}

	// res % (n+1) 是为了处理 sum(a) < s 的情况
	return res % (n + 1)
}
```



```python
class Solution(object):
    def minSubArrayLen(self, s, nums):
        """
        :type s: int
        :type nums: List[int]
        :rtype: int
        """
        start = 0
        sum_of_subarray = 0
        min_len = float("inf")
        for end in range(len(nums)):
            sum_of_subarray += nums[end]
            while sum_of_subarray >= s:
                min_len = min(min_len, end-start+1)
                sum_of_subarray -= nums[start]
                start += 1
        return min_len if min_len != float("inf") else 0
```