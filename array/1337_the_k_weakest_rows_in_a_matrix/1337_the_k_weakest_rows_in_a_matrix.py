
class Solution(object):
    def kWeakestRows(self, mat, k):
        l = [[sum(row), i] for i, row in enumerate(mat)]
        s = sorted(l)
        return [r[1] for r in s[:k]]


mat = [[1,1,0,0,0],[1,1,1,1,0],[1,0,0,0,0],[1,1,0,0,0],[1,1,1,1,1]]
k = 3
res = Solution().kWeakestRows(mat, k)
print(res)