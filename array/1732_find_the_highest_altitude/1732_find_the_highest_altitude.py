from typing import List

class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        a = [0]
        for i in range(len(gain)):
            a.append(a[i]+gain[i])
        return max(a)


gain = [-5,1,5,0,-7]
res = Solution().largestAltitude(gain)
print(res)