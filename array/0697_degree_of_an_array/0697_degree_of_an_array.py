# -*- coding: utf-8 -*-



class Solution(object):
    def findShortestSubArray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        first, counter, res, degree = {}, {}, 0, 0
        for i, v in enumerate(nums):
            first.setdefault(v, i)
            counter[v] = counter.get(v, 0) + 1
            if counter[v] > degree:
                degree = counter[v]
                res = i - first[v] + 1
            elif counter[v] == degree:
                res = min(res, i - first[v] + 1)
        return res



print(Solution().findShortestSubArray([1,2,2,3,1]))