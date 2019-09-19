class Solution(object):
    def rotate(self, matrix):
        n = len(matrix)

        for i in xrange(n):
            for j in xrange(n - i):
                matrix[i][j], matrix[n-1-j][n-1-i] = matrix[n-1-j][n-1-i], matrix[i][j]

        for i in xrange(n/2):
            for j in xrange(n):
                matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
        
        return matrix


matrix = [[1,2,3],[4,5,6],[7,8,9]] 

res = Solution().rotate(matrix)
print(res)