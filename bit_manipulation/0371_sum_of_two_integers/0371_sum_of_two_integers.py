# -*- coding: utf-8 -*-

class Solution(object):
    def getSum(self, a, b):
        """
        :type a: int
        :type b: int
        :rtype: int
        """
        return ((a & b) << 1) + (a ^ b)

print(Solution().getSum(5,7))