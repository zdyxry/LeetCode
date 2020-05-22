from typing import List

class Solution:
    def isMonotonic(self, A: List[int]) -> bool:
        cnt_inc = 0; cnt_dec = 0
        for i, v in enumerate(A):
            if i == 0: prev = v; continue
            if   v > prev: cnt_inc += 1
            elif v < prev: cnt_dec += 1
            if cnt_inc and cnt_dec: return False
            prev = v
        return True


A = [1,2,2,3]
res = Solution().isMonotonic(A)
print(res)