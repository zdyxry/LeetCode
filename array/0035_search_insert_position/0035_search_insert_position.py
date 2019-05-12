# -*- coding: utf-8 -*-
class Solution(object):
    def searchInsert(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = left + (right - left) /2
            if nums[mid] >= target:
                right = mid - 1
            else:
                left = mid + 1
                
        return left


nums = [1,2,3,4,6]
target = 5

print(Solution().searchInsert(nums, target))