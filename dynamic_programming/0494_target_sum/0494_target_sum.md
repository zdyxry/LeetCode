494. Target Sum


Medium


You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols + and -. For each integer, you should choose one from + and - as its new symbol.

Find out how many ways to assign symbols to make sum of integers equal to target S.

Example 1:
```
Input: nums is [1, 1, 1, 1, 1], S is 3. 
Output: 5
Explanation: 

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

There are 5 ways to assign symbols to make the sum of nums be target 3.
```

Note:

The length of the given array is positive and will not exceed 20.

The sum of elements in the given array will not exceed 1000.

Your output answer is guaranteed to be fitted in a 32-bit integer.


## 方法

```go
func findTargetSumWays(nums []int, S int) int {
    sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum < S {
		return 0
	}
	// sum 和 S 必须同为奇数或者偶数
	if (sum+S)%2 == 1 {
		return 0
	}

	return calcSumWays(nums, (sum+S)/2)
}

// nums 被分为两个部分 P 和 N
// P 中的元素前面放的是 +
// N 中的元素前面放的是 -
// 可得
//                   sum(P) - sum(N) = target
// sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
//                        2 * sum(P) = target + sum(nums)
// 于是，题目就被转换成了
// nums 有多少个子集，其和为 (target + sum(nums))/2
func calcSumWays(nums []int, target int) int {
	// dp[i] 表示 nums 中和为 i 的子集个数
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target]
}
```


```python
class Solution(object):
    def findTargetSumWays(self, nums, S):
        """
        :type nums: List[int]
        :type S: int
        :rtype: int
        """
        from collections import defaultdict
        memo = {0: 1}
        for x in nums:
            m = defaultdict(int)
            for s, n in memo.iteritems():
                m[s + x] += n
                m[s - x] += n
            memo = m
        return memo[S]
```