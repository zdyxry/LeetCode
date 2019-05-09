# -*- coding: utf-8 -*-

class Solution(object):
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        idx, cnt = 0, 1

        for i in xrange(1, len(nums)):
            if nums[idx] == nums[i]:
                cnt += 1
            else:
                cnt -= 1
                if cnt == 0:
                    idx = i
                    cnt = 1
        return nums[idx]

    def majorityElement2(self, nums):
        elem = count = 0 
        for i in nums:
            if count == 0:
                elem = i
                count = 1
            else:
                if elem == i:
                    count += 1
                else:
                    count -=1
        return elem

    def majorityElement3(self, nums):
        my_dict = {}
        for i in nums:
            if not my_dict.get(i):
                my_dict[i] = 1
            else:
                my_dict[i] += 1
            if my_dict[i] > len(nums)/2:
                return i
        


a = [1,2,2,2,3,3,3,3,3,3,4]
print(Solution().majorityElement3(a))