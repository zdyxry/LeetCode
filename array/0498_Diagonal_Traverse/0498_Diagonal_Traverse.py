import collections
from typing import List


class Solution:
    def findDiagonalOrder(self, matrix: List[List[int]]) -> List[int]:
        map = {}
        res = []
        for r, _ in enumerate(matrix):
            for c, _ in enumerate(matrix[r]):
                if not map.get(r+c):
                    map[r+c] = []
                map[r+c].append(matrix[r][c])
        od = collections.OrderedDict(sorted(map.items()))
        flag = False
        for k, v in od.items():
            if flag:
                res.extend(v)
                flag = False
            else:
                res.extend(reversed(v))
                flag = True
        return res


matrix = [[1,2,3],[4,5,6],[7,8,9]]
res = Solution().findDiagonalOrder(matrix)
print(res)