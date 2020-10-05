from typing import List


class Solution:
    def restoreMatrix(self, rowSum: List[int], colSum: List[int]) -> List[List[int]]:
        m=len(rowSum)
        n=len(colSum)
        
        res=[[0]*n for _ in range(m)]
        for i in range(m):
            for j in range(n):
                if rowSum[i]>=colSum[j]:
                    res[i][j] = colSum[j]
                    rowSum[i] = rowSum[i]-colSum[j]
                    colSum[j]=0
                else:
                    res[i][j] = rowSum[i]
                    colSum[j] = colSum[j]-rowSum[i]
                    rowSum[i] = 0
                    
                # print(rowSum,colSum)
        return res