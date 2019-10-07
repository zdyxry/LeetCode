class Solution(object):
    def minimumTotal(self, triangle):
        n = len(triangle)
        if n == 0:
            return None 
        if n == 1:
            return triangle[0][0]
        for i in range(1, n):
            triangle[i][0] += triangle[i-1][0]
            triangle[i][-1] += triangle[i-1][-1]
            for j in range(1, i):
                triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
        return min(triangle[-1])

triangle = [
    [2],
    [3,4],
    [6,5,7],
    [4,1,8,3]
]

res = Solution().minimumTotal(triangle)
print(res)