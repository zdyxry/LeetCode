from typing import List


class Solution:
    def countGoodRectangles(self, rectangles: List[List[int]]) -> int:
        res = 0
        maxlen = max([min(i) for i in rectangles])
        for j in rectangles:
            if maxlen == min(j):
                res += 1
        return res


rectangles = [[5,8],[3,9],[5,12],[16,5]]
res = Solution().countGoodRectangles(rectangles)
print(res)