# -*- coding: utf-8 -*-


class Solution(object):
    def findDisappearedNumbers(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        for i in xrange(len(nums)):
            if nums[abs(nums[i]) - 1] > 0:
                nums[abs(nums[i]) - 1] *= -1

        result = []
        for i in xrange(len(nums)):
            if nums[i] > 0:
                result.append(i+1)
            else:
                nums[i] *= -1
        return result


nums = [4,3,2,7,8,2,3,1] 
print(Solution().findDisappearedNumbers(nums))