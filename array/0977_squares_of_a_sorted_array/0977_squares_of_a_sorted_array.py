# -*- coding: utf-8 -*-



class Solution(object):
    def sortedSquares(self, A):
        """
        :type A: List[int]
        :rtype: List[int]
        """

        size = len(A)
        res = [0] * size

        l, r, i = 0, size - 1, size - 1
        while l <= r:
            if A[l] + A[r] < 0:
                res[i] = A[l] * A[l]
                l += 1
            else:
                res[i] = A[r] * A[r]
                r -= 1
            i -= 1
        return res


A = [-4,-1,0,3,10]
print(Solution().sortedSquares(A))