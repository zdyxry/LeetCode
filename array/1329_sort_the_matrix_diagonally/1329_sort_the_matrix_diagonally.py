from collections import defaultdict


class Solution(object):
    def diagonalSort(self, mat):
        diag = defaultdict(list)
        N, M = len(mat), len(mat[0])

        for i in range(N):
            for j in range(M):
                diag[i-j].append(mat[i][j])
        print diag
        for k in diag.keys():
            diag[k].sort(reverse=True)
        print diag
        for i in range(N):
            for j in range(M):
                mat[i][j] = diag[i-j].pop()

        return mat



mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]

res = Solution().diagonalSort(mat)
print(res)