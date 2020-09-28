from typing import List

class Solution:
    def flipAndInvertImage(self, A: List[List[int]]) -> List[List[int]]:
        m , n = len(A), len(A[0])
        for i in range(m):
            j, k = 0, n -1
            while j < k:
                A[i][j], A[i][k] = A[i][k], A[i][j]
                j += 1
                k -= 1
        
        for i in range(m):
            for j in range(n):
                A[i][j] = 1 if A[i][j] == 0 else 0
        
        return A