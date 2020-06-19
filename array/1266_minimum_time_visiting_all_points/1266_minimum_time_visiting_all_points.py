from typing import List

class Solution:
    def minTimeToVisitAllPoints(self, points: List[List[int]]) -> int:
        last_point=points[0]
        step=0
        for point in points:
            x=abs(last_point[0]-point[0])
            y=abs(last_point[1]-point[1])
            last_point=point
            if x>=y:
                step+=x
            else:
                step+=y
        return step


points = [[1,1],[3,4],[-1,0]]
res = Solution().minTimeToVisitAllPoints(points)
print(res)