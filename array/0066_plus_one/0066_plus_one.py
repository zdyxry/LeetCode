# -*- coding: utf-8 -*-

class Solution(object):
    def plusOne(self, digits):
        """
        :type digits: List[int]
        :rtype: List[int]
        """
        old = ""
        for i in digits:
            old += str(i)
        old = int(old)
        
        result = []
        new = old + 1
        for j in str(new):
            result.append(int(j))
        return result

print(Solution().plusOne([1,2,3]))