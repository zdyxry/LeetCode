# -*- coding: utf-8 -*-

class Solution(object):
    def canPlaceFlowers(self, flowerbed, n):
        """
        :type flowerbed: List[int]
        :type n: int
        :rtype: bool
        """
        size = len(flowerbed)
        for i in xrange(size):
            if flowerbed[i] == 0:
                avail = True
                if i - 1 >= 0 and flowerbed[i-1]==1:
                    avail = False
                if i + 1 < size and flowerbed[i+1]==1:
                    avail = False
                if avail:
                    flowerbed[i] = 1
                    n -= 1
        return n <= 0



print(Solution().canPlaceFlowers([1,0,0,0,0,1], 2))