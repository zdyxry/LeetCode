class Solution(object):
    def largestPerimeter(self, A):
        A.sort()
        for i in reversed(xrange(len(A)- 2)):
            if A[i] + A[i+1] > A[i+2]:
                return A[i] + A[i+1] + A[i+2]
        return 0


A = [1,2,1]
res = Solution().largestPerimeter(A)

print(res)