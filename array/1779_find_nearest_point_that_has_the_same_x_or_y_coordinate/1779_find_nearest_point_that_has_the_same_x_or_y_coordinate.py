from typing import List


class Solution:
    def nearestValidPoint(self, x: int, y: int, points: List[List[int]]) -> int:
        m = float('INF')
        res = -1
        for idx, p in enumerate(points):
            if p[0] == x or p[1] == y:
                if abs(p[0]-x) + abs(p[1]-y) < m:
                    m = abs(p[0]-x) + abs(p[1]-y)
                    res = idx
        return res
                
