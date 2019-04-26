# -*- coding: utf-8 -*-

class Solution(object):
    def hammingDistance(self, x, y):
        """
        :type x: int
        :type y: int
        :rtype: int
        """
        x ^= y 
        
        res = 0 
        while x:
            res += x & 1
            x >>= 1
        return res

print(Solution().hammingDistance(1,5))