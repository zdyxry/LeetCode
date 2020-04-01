
class Solution:
    def maxSumDivThree(self, nums):
        dp = [0, 0, 0]
        
        for n in nums:
            tmp_dp = dp[:]
            for i in range(len(dp)):
                c_sum = tmp_dp[i] + n
                dp[c_sum % 3] = max(dp[c_sum % 3], c_sum)

        return dp[0]


nums = [3,6,5,1,8]
res = Solution().maxSumDivThree(nums)
print(res)