from typing import List

class Solution:
    def repeatedNTimes(self, A: List[int]) -> int:
        for k in range(1, 4):
            for i in range(len(A) - k):
                if A[i] == A[i+k]:
                    return A[i]


A = [1,2,3,3]
res = Solution().repeatedNTimes(A)
print(res)