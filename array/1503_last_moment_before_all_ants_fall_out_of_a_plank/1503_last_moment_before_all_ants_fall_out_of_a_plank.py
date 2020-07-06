from typing import List


class Solution:
    def getLastMoment(self, n: int, left: List[int], right: List[int]) -> int:
        time = 0
        if left:
            time = max(time, max(left))
        if right:
            time = max(time, n - min(right))
        return time



n = 7
left = []
right = [0,1,2,3,4,5,6,7]
res = Solution().getLastMoment(n, left, right)
print(res)