from typing import List


class Solution:
    def queensAttacktheKing(self, queens: List[List[int]], king: List[int]) -> List[List[int]]:
        directions = [[0,1],[1,0],[1,1],[-1,0],[0,-1],[-1,-1],[1,-1],[-1,1]]
        x0,y0 = king
        queens = set(map(tuple,queens))
        res = []
        for x,y in directions:
            x1,y1 = x0+x, y0+y
            for i in range(8):
                if (x1 >= 0 and x1 <= 7) and (y1 >= 0 and y1 <= 7):
                    if (x1,y1) in queens:
                        res.append([x1,y1])
                        break
                x1 += x
                y1 += y
        return res