# -*- coding: utf-8 -*-


class Solution(object):

    def findLengthOfLCIS(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        result, count = 0, 0
        for i in xrange(len(nums)):
            if i == 0 or nums[i - 1] < nums[i]:
                count += 1
                result = max(result, count)
            else:
                count = 1
        return result


print(Solution().findLengthOfLCIS([1, 3, 5, 4, 7]))