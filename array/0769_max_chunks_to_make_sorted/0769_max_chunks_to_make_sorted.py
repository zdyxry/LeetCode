from typing import List


class Solution:
    def maxChunksToSorted(self, arr: List[int]) -> int:
        A , B = 0 , 0
        for i,j in enumerate(arr):
            A = max(A,j)
            if(A == i):B += 1
        return B 



arr = [4,3,2,1,0]
res = Solution().maxChunksToSorted(arr)
print(res)