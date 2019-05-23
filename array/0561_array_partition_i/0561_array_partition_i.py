# -*- coding: utf-8 -*-

class Solution(object):
    def arrayPairSum(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        nums.sort()
        result = 0
        for i in xrange(0, len(nums), 2):
            result += nums[i]
        return result


nums = [1,4,3,2]
print(Solution().arrayPairSum(nums))