# -*- coding: utf-8 -*-


class Solution(object):
    def maxSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if max(nums) < 0:
            return max(nums)
        
        global_max, local_max = float('-inf'), 0
        for num in nums:
            local_max = max(0, local_max + num)
            global_max = max(local_max, global_max)
        return global_max



nums = [-2,1,-3,4,-1,2,1,-5,4]
print(Solution().maxSubArray(nums))