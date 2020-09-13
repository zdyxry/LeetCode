from typing import List

class Solution:
    def numSpecial(self, mat: List[List[int]]) -> int:
        m, n, ans = len(mat), len(mat[0]), 0
        row, col = [0] * m, [0] * n
        for i in range(m):
            for j in range(n):
                if mat[i][j]:
                    row[i] += 1
                    col[j] += 1
        pool = [j for j in range(n) if col[j] == 1]
        for i in range(m):
            if row[i] != 1:
                continue
            for j in pool:
                if mat[i][j]:
                    ans += 1
                    break
        return ans


mat = [[1,0,0],[0,0,1],[1,0,0]]
res = Solution().numSpecial(mat)
print(res)