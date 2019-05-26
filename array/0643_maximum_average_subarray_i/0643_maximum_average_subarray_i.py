# -*- coding: utf-8 -*-

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


print(Solution().findMaxAverage([1,12,-5,-6,50,3], 4))