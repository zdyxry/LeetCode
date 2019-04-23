# -*- coding: utf-8 -*-

class Solution:
    def singleNumber(self, nums):
        import operator
        from functools import reduce
        return reduce(operator.xor, nums)

nums = [1,2,2,1,3]
print(Solution().singleNumber(nums))