from typing import List

class Solution:
    def prefixesDivBy5(self, A: List[int]) -> List[bool]:
        result = []
        for i in range(len(A)):
            if i: 
                A[i] = A[i] + (A[i-1] << 1)
            if A[i] % 5 == 0: 
                result.append(True)
            else:
                result.append(False)
        return result  



A = [0,1,1]
res = Solution().prefixesDivBy5(A)
print(res)