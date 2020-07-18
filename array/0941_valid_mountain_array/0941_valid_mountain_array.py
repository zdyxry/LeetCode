from typing import List


class Solution:
    def validMountainArray(self, A: List[int]) -> bool:
        l,r=0,len(A)-1
        while l<r and A[l]<A[l+1]: l+=1
        while r>l and A[r]<A[r-1]: r-=1
        return l==r and l!=0 and r!=len(A)-1


A = [2,1]
res = Solution().validMountainArray(A)
print(res)