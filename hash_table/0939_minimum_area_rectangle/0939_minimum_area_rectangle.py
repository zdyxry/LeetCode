from typing import List

class Solution:
    def minAreaRect(self, points: List[List[int]]) -> int:
        area = 0
        mem = set()
        for x1,y1 in points:
            for x2,y2 in mem:
                if (x2,y1) in mem and (x1,y2) in mem:
                    if not area:
                        area = abs(x2-x1)*abs(y2-y1)
                    else:
                        if abs(x2-x1)*abs(y2-y1) < area:
                            area = abs(x2-x1)*abs(y2-y1)
            mem.add((x1,y1))
        return area



points = [[1,1],[1,3],[3,1],[3,3],[2,2]]
res = Solution().minAreaRect(points)
print(res)