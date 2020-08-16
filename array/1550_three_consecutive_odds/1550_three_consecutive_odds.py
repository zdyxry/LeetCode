from typing import List

class Solution:
    def threeConsecutiveOdds(self, arr: List[int]) -> bool:
        cnt = 0
        for i in arr:
            if i & 1:
                cnt += 1
            else:
                cnt = 0
            if cnt == 3:
                return True
        return False


arr = [2,6,4,1]
res = Solution().threeConsecutiveOdds(arr)
print(res)