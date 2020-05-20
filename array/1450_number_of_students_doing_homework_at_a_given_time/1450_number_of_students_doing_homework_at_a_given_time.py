from typing import List

class Solution:
    def busyStudent(self, startTime: List[int], endTime: List[int], queryTime: int) -> int:
        return sum(a<=queryTime<=b for a,b in zip(startTime,endTime))


startTime = [1,2,3]
endTime = [3,2,7]
queryTime = 4
res = Solution().busyStudent(startTime, endTime, queryTime)
print(res)