from typing import List


class Solution:
    def diagonalSum(self, mat: List[List[int]]) -> int:
        if len(mat) == 1:
            return mat[0][0]
        n = len(mat)
        ans = 0

        for i in range(n):
            if i != n - 1 -i:
                ans += mat[i][i]
                ans += mat[i][n-1-i]
            else:
                ans += mat[i][i]
        return ans


mat = [[1,2,3],[4,5,6],[7,8,9]]
res = Solution().diagonalSum(mat)
print(res)