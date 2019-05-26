643. Maximum Average Subarray I

Easy

Given an array consisting of n integers, find the contiguous subarray of given length k that has the maximum average value. And you need to output the maximum average value.

Example 1:
```
Input: [1,12,-5,-6,50,3], k = 4
Output: 12.75
Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
 ```

Note:

1 <= k <= n <= 30,000.

Elements of the given array will be in the range [-10,000, 10,000].

# 方法
遍历数组，求出长度为 k 时最大子数组之和。


```go
func findMaxAverage(nums []int, k int) float64 {
    total := 0
    for i:=0;i<k;i++ {
        total += nums[i]
    }
    
    result := total
    
    for i:=k;i<len(nums);i++ {
        total = total + nums[i] - nums[i-k]
        if result < total {
            result = total
        }
    }
    return float64(result)/float64(k)
}
```


```python
class Solution(object):
    def findMaxAverage(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: float
        """
        result = total = sum(nums[:k])
        for i in xrange(k, len(nums)):
            total = total + nums[i] - nums[i-k]
            result = max(total, result)
            
        return float(result)/k
            
```