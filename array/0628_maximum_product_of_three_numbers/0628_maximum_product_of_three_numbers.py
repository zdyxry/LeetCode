# -*- coding: utf-8 -*-

class Solution(object):
    def maximumProduct(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        min1, min2 = float('inf'), float('inf')
        max1, max2, max3 = float('-inf'), float('-inf'), float('-inf')
        
        for num in nums:
            if num > max1:
                max3 = max2
                max2 = max1
                max1 = num
            elif num > max2:
                max3 = max2
                max2 = num
            elif num > max3:
                max3 = num
            
            if num < min1:
                min2 = min1
                min1 = num
            elif num < min2:
                min2 = num
            
        return max(max3*max2, min1*min2) * max1


print(Solution().maximumProduct([1,2,3]))