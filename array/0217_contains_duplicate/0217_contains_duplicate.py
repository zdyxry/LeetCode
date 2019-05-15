# -*- coding: utf-8 -*-


class Solution(object):
    def containsDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        d = {}
        for n in nums:
            if n in d:
                return True
            d[n] = True
        return False

nums = [1,2,3,1]
print(Solution().containsDuplicate(nums))