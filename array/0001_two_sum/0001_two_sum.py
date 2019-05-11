# -*- coding: utf-8 -*-
class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        d = {}
        for idx, num in enumerate(nums):
            if (target - num) in d:
                return [d[target-num], idx]
            d[num] = idx


nums = [2,7,11,15]
target = 9
print(Solution().twoSum(nums, target))