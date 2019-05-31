724. Find Pivot Index

Easy

Given an array of integers nums, write a method that returns the "pivot" index of this array.

We define the pivot index as the index where the sum of the numbers to the left of the index is equal to the sum of the numbers to the right of the index.

If no such index exists, we should return -1. If there are multiple pivot indexes, you should return the left-most pivot index.

Example 1:

```
Input: 
nums = [1, 7, 3, 6, 5, 6]
Output: 3
Explanation: 
The sum of the numbers to the left of index 3 (nums[3] = 6) is equal to the sum of numbers to the right of index 3.
Also, 3 is the first index where this occurs.
```
 

Example 2:

```
Input: 
nums = [1, 2, 3]
Output: -1
Explanation: 
There is no index that satisfies the conditions in the problem statement.
```
 

Note:

The length of nums will be in the range [0, 10000].

Each element nums[i] will be an integer in the range [-1000, 1000].




# 方法
两次遍历，第一次先计算所有数字之和。
第二次，当数字之和减去当前值等于 2*当前值左边所有数值之和，则规则成立，因为是从前向后遍历，因此返回第一个规则成立的数字索引即可。



```go
func pivotIndex(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    tsum := 0
    for i, num := range nums {
        if (sum-num) == 2*tsum {
            return i
        }
        tsum += num
    }
    return -1
}
```


```python
class Solution(object):
    def pivotIndex(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        total = sum(nums)
        left_sum = 0
        for i, num in enumerate(nums):
            if left_sum == (total-left_sum-num):
                return i
            left_sum += num
        return -1
```