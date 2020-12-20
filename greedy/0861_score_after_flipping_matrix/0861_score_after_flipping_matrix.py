from typing import List

class Solution:
    def matrixScore(self, A: List[List[int]]) -> int:
        row = len(A)
        col = len(A[0])



        #先把所有的第一位都变成1
        for i in range(row):
            if A[i][0]==0:
                for j in range(len(A[i])):
                    A[i][j]=1-A[i][j]
        #后面的列，0多的就转，1多的就保留
        for i in range(1,col):
            num0 = 0
            num1 = 0
            #统计个数
            for k in range(row):
                if A[k][i]==0:
                    num0 +=1
                else:
                    num1 +=1
            if num1<num0:
                for j in range(row):
                    A[j][i]=1-A[j][i]

        res = 0
        for i in range(len(A)):
            temp = ""
            for j in range(len(A[0])):
                temp+=str(A[i][j])
            res+=int(temp,2)
        return res 