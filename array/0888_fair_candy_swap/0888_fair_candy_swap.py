from typing import List


class Solution:
    def fairCandySwap(self, A: List[int], B: List[int]) -> List[int]:
        Sa, Sb = sum(A), sum(B)
        setB = set(B)
        for x in A:
            if x + (Sb - Sa) / 2 in setB:
                return [x, x+(Sb-Sa) / 2]


A = [1,1]
B = [2,2]
res = Solution().fairCandySwap(A, B)
print(res)