from typing import List

class Solution:
    def findSmallestSetOfVertices(self, n: int, edges: List[List[int]]) -> List[int]:
        a = set()
        b = set()
        for x, y in edges:
            a.add(x)
            b.add(y)
        return list(a - b)


n = 6
edges = [[0,1],[0,2],[2,5],[3,4],[4,2]]
res = Solution().findSmallestSetOfVertices(n, edges)
print(res)