class Solution(object):
    def minSubArrayLen(self, s, nums):
        start = 0
        sum_of_subarray = 0
        min_len = float('inf')

        for end in range(len(nums)):
            sum_of_subarray += nums[end]
            while sum_of_subarray >= s:
                min_len = min(min_len, end-start+1)
                sum_of_subarray -= nums[start]
                start += 1
        return min_len if min_len != float('inf') else 0


nums = [2,3,1,2,4,3]
s = 7
res = Solution().minSubArrayLen(s, nums)
print(res)