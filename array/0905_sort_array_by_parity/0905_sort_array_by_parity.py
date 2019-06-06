# -*- coding: utf-8 -*-
class Solution(object):
    def sortArrayByParity(self, A):
        """
        :type A: List[int]
        :rtype: List[int]
        """
        if not A:
            return []
        
        start, end = 0, len(A)-1
        
        while (start<end):
            while (start<end and A[start]%2==0):
                start+=1
            while (start<end and A[end]%2==1):
                end-=1
            
            if start<end:
                A[start],A[end]=A[end],A[start]
                start+=1
                end-=1
        
        return A


print(Solution().sortArrayByParity(
[3,1,2,4]))