from typing import List


class Solution:
    def maxAscendingSum(self, nums: List[int]) -> int:
        max_sum = nums[0]
        s = nums[0]
        for i in range(1, len(nums)):
            if nums[i] > nums[i-1]:
                s += nums[i]
            else:
                s = nums[i]
            if max_sum < s:
                max_sum = s
        return max_sum



nums = [10,20,30,5,10,50]
res = Solution().maxAscendingSum(nums)
print(res)