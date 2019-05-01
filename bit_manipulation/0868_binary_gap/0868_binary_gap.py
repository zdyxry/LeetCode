# -*- coding: utf-8 -*-

class Solution(object):
    def binaryGap(self, N):
        """
        :type N: int
        :rtype: int
        """
        pre = dist = 0
        for i, c in enumerate(bin(N)[2:]):
            if c == "1":
                dist = max(dist, i - pre)
                pre = i
        return dist

print(Solution().binaryGap(5))