# -*- coding: utf-8 -*-


class Solution(object):
    def missingNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        import operator
        return reduce(operator.xor, nums, \
                      reduce(operator.xor, xrange(len(nums) + 1)))


print(Solution().missingNumber([1,2,3,4,5,6]))