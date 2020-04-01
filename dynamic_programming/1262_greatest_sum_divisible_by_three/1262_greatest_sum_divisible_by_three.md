1262. Greatest Sum Divisible by Three


Medium


Given an array nums of integers, we need to find the maximum possible sum of elements of the array such that it is divisible by three.

 

Example 1:

```
Input: nums = [3,6,5,1,8]
Output: 18
Explanation: Pick numbers 3, 6, 1 and 8 their sum is 18 (maximum sum divisible by 3).
```

Example 2:

```
Input: nums = [4]
Output: 0
Explanation: Since 4 is not divisible by 3, do not pick any number.
```

Example 3:

```
Input: nums = [1,2,3,4,4]
Output: 12
Explanation: Pick numbers 1, 3, 4 and 4 their sum is 12 (maximum sum divisible by 3).
```

Constraints:

1 <= nums.length <= 4 * 10^4

1 <= nums[i] <= 10^4


## 方法

```go
func maxSumDivThree(nums []int) int {
    res := 0
    one := 10000
    two := 10000
    for i := 0; i < len(nums); i++ {
        res += nums[i]
        if nums[i] % 3 == 1 {
            two = min(two, one + nums[i])
            one = min(one, nums[i])
        }
        if nums[i] % 3 == 2 {
            one = min(one, two + nums[i])
            two = min(two, nums[i])
        }
    }
    if res % 3 == 0 { return res }
    if res % 3 == 1 { return res - one }
    return res - two
}

func min(a, b int) int {
    if a < b { return a }
    return b
}
```



```python
class Solution:
    def maxSumDivThree(self, nums: List[int]) -> int:
        dp = [0, 0, 0]
        
        for n in nums:
            tmp_dp = dp[:]
            for i in range(len(dp)):
                c_sum = tmp_dp[i] + n
                dp[c_sum % 3] = max(dp[c_sum % 3], c_sum)


        return dp[0]
```