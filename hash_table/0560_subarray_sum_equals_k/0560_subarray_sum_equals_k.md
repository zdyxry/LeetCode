560. Subarray Sum Equals K


Medium


Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.

Example 1:
```
Input:nums = [1,1,1], k = 2
Output: 2
```


Note:
The length of the array is in range [1, 20,000].

The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].

## 方法

```go
func subarraySum(a []int, k int) int {
    res, sum := 0, 0
	rec := make(map[int]int, len(a))
	rec[0] = 1

	for i := range a {
		// sum 等于 a[:i] 中所有元素之和
		sum += a[i]
		// rec[x] 是 a[:i] 中所有和为 x 的连续子序列的个数
		// 假设 rec[sum-k] == 1
		// 那么，有且仅有一个 j，满足 0<= j < i，且 a[j:i] 中所有元素的和为 k
		res += rec[sum-k]
		// 更新 rec
		rec[sum]++
	}

	return res
}
```


```python
class Solution(object):
    def subarraySum(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        count,p,d = 0,0,{0:1}
        for i in nums:
            p+=i
            if p-k in d:
                count+=d[p-k]
            if p not in d:
                d[p] = 1
            else:
                d[p] +=1
        return count
```