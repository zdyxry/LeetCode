1695. Maximum Erasure Value


Medium


You are given an array of positive integers nums and want to erase a subarray containing unique elements. The score you get by erasing the subarray is equal to the sum of its elements.

Return the maximum score you can get by erasing exactly one subarray.

An array b is called to be a subarray of a if it forms a contiguous subsequence of a, that is, if it is equal to a[l],a[l+1],...,a[r] for some (l,r).

 

Example 1:

```
Input: nums = [4,2,4,5,6]
Output: 17
Explanation: The optimal subarray here is [2,4,5,6].
```

Example 2:

```
Input: nums = [5,2,1,2,5,2,1,2,5]
Output: 8
Explanation: The optimal subarray here is [5,2,1] or [1,2,5].
```

Constraints:

1 <= nums.length <= 105    
1 <= nums[i] <= 104


## 方法

```go
func maximumUniqueSubarray(nums []int) int {

    count := make([]int, 10001)
    left := 0
    res := nums[0]
    sum := nums[0]
    count[nums[0]]++
    for i := 1; i < len(nums); i++ {
        sum += nums[i]
        count[nums[i]]++
        for count[nums[i]] > 1 {
            count[nums[left]]--
            sum -= nums[left]
            left++
        }
        res = max(res, sum)
    }

    return res
}


func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

```


```python
class Solution:
    def maximumUniqueSubarray(self, nums: List[int]) -> int:
        dic = {}  
        j = 0
        res = 0
        s = 0
        length = len(nums)
        for i in range(length):
            if nums[i] not in dic:
                dic[nums[i]] = 1
            else:
                dic[nums[i]] += 1
            s += nums[i]
            while dic[nums[i]] > 1:
                dic[nums[j]] -= 1
                s -= nums[j]
                j += 1
            res = max(res,s)
        return res if res > s else s
```