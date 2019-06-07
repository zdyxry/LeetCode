# -*- coding: utf-8 -*-

class Solution(object):
    def sortArrayByParityII(self, A):
        """
        :type A: List[int]
        :rtype: List[int]
        """
        j = 1
        for i in xrange(0, len(A), 2):
            if A[i] % 2:
                while A[j] % 2:
                    j += 2
                A[i], A[j] = A[j], A[i]
            print(A)
        return A

A = [4,2,5,7]
print(Solution().sortArrayByParityII(A))