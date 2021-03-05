from typing import List

class Solution:
    def matrixBlockSum(self, mat: List[List[int]], K: int) -> List[List[int]]:
        m, n = len(mat), len(mat[0])
        prefix_sum = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                prefix_sum[i][j] = prefix_sum[i - 1][j] + 
                                   prefix_sum[i][j - 1] - 
                                   prefix_sum[i - 1][j - 1] + 
                                   mat[i - 1][j - 1]

        res = [[0] * n for _ in range(m)]
        for i in range(m):
            for j in range(n):
                min_row = i - K if i - K >= 0 else 0
                max_row = i + K if i + K < m else m - 1
                min_col = j - K if j - K >= 0 else 0
                max_col = j + K if j + K < n else n - 1
                res[i][j] = prefix_sum[max_row + 1][max_col + 1] - 
                            prefix_sum[max_row + 1][min_col] - 
                            prefix_sum[min_row][max_col + 1] + 
                            prefix_sum[min_row][min_col]
        return res


mat = [[1,2,3],[4,5,6],[7,8,9]]
K = 1
res = Solution().matrixBlockSum(mat, K)
print(res)