from typing import List

class Solution:
    def checkStraightLine(self, coordinates: List[List[int]]) -> bool:
        same_slop = True
        last_slop = None
        intial_point = coordinates[0]
        for point in coordinates[1:]:
            try:
                slop = (point[1] - intial_point[1])/(point[0] - intial_point[0])
            except ZeroDivisionError: 
                slop = float(inf)
            if last_slop== None:last_slop = slop
            elif slop == last_slop:continue
            else: 
                same_slop = False
                break
                
        return same_slop


coordinates =[[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
res = Solution().checkStraightLine(coordinates)
print(res)