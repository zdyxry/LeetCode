class Solution(object):
    def findLength(self, A, B):
        if len(A) < len(B): return self.findLength(B, A)
        result = 0
        dp = [[0] * (len(B)+1) for _ in xrange(2)]

        for i in xrange(len(A)):
            for j in xrange(len(B)):
                if A[i] == B[j]:
                    dp[(i+1)%2][j+1] = dp[i%2][j] + 1
                else:
                    dp[(i+1)%2][j+1] = 0
            result = max(result, max(dp[(i+1)%2]))
        return result


A = [1,2,3,2,1]
B = [3,2,1,4,7]

res = Solution().findLength(A, B)
print(res)